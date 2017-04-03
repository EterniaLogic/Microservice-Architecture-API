package Comment 

// User: Brent Clancy (EterniaLogic)
// Date: 1/25/2015

import (
    "net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"../../common"
	"strconv"
)


// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/comments";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
		
	// Get/Modify
	router.HandleFunc(prefix+"/{vid}", HandleGetVidComments).Methods("GET");	// Get list of comments for a specific video ID
	
	router.HandleFunc(prefix+"", HandlePostComment).Methods("POST");	// Get list of comments for a specific video ID
		
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}


func HandleGetVidComments(w http.ResponseWriter, r *http.Request){
	// Handle Custom JSON for a search
	id := mux.Vars(r)["vid"];
	
	vid,_ := strconv.ParseInt(id,10,64);
	
	// get SQL data
	rows, _ := common.GetDB().Query("SELECT HEX(UUID),`comment`,`date` FROM `"+common.GetConfig().DBTable+"` WHERE `VID`=? ORDER BY `date` DESC LIMIT 20",vid);
	

	
	// output data
	w.Header().Set("Content-Type","application/json");
	w.Write([]byte(GenJSONRows(rows)));
}

func HandlePostComment(w http.ResponseWriter, r *http.Request){
	// Handle Custom JSON for a search
	returnstring := "";
	
	
	ContentType := r.Header["Content-Type"];
	
	// get the Token from the header
	if(len(ContentType) > 0){
		if(common.RESTFulTestToken(r)){
			if(common.RESTFulGetToken(r) != ""){
				UUID := common.GetIDFromToken(common.RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && UUID != ""){
					data,_ := ioutil.ReadAll(r.Body);
					txt,err := common.JSONtoMap(string(data));
					
					if(err){
						// JSON error
						common.AsyncPrintln("[Error](HandlePostComment) JSON error");
						
						w.Header().Set("Content-Type","application/json");
						w.Write(common.JSONPackageMessage("","Bad JSON formatting",false));
						
					}else{
						// JSON transfered over, inspect the output
						//returnstring = common.TestJSONInput(txt["VID"],"JSON VID not returned!").(string);
						//returnstring = common.TestJSONInput(txt["Comment"],"JSON Comment not returned!").(string);
						
						if(returnstring == ""){
							// Post the comment for the video
							vidtt := int64(txt["VID"].(float64));
							PostComment(w,r, vidtt, txt["Comment"].(string));
							returnstring = "{\"Type\":\"Success\", \"Message\":\"Comment created\"}";
						}
						
						w.Header().Set("Content-Type","application/json");
						w.Write([]byte(returnstring));
					}
				}else{
					w.Header().Set("Content-Type","application/json");
					w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Unknown token 2 or bad Content Type\"}"));
				}
			}else{
				w.Header().Set("Content-Type","application/json");
				w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Unknown Token\"}"));
			}
		}else{
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"No token provided\"}"));
		}
	}else{
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Not 'application/json' MIME\"}"));
	}
}