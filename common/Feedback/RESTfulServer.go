package Feedback 

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "net/http"
	"github.com/gorilla/mux"
	"../../common"
	"io/ioutil"
)


// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/feedback";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
		
	// Add new
	router.HandleFunc(prefix, NewFeedback).Methods("POST");
	
	// return # of feedbacks
	//router.HandleFunc(prefix+"/recent/{num}", GetRecentFeedback).Methods("GET"); // return the recent {num} of feedbacks
		
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}

func NewFeedback(w http.ResponseWriter, r *http.Request) {
	ContentType := r.Header["Content-Type"];
	
	// get the Token from the header
	if(len(ContentType) > 0){
		if(common.RESTFulTestToken(r)){
			if(common.RESTFulGetToken(r) != ""){
				UUID := common.GetIDFromToken(common.RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && UUID != ""){
					var responseJSON interface{};
					body,_:=ioutil.ReadAll(r.Body);
					smap,err := common.JSONtoMap(string(body));
					if(!err){
						responseJSON = common.TestJSONInput(smap["Type"], "JSON `Type` key not found");
						responseJSON = common.TestJSONInput(smap["Title"], "JSON `Title` key not found");
						responseJSON = common.TestJSONInput(smap["Action"], "JSON `Action` key not found");
						responseJSON = common.TestJSONInput(smap["Description"], "JSON `Description` key not found");	
						
						if(responseJSON == nil){
							//msg,err := AuthRegister(smap["Username"].(string), smap["UserPass"].(string), smap["Email"].(string));
							err,_ := common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"` (UUID,type,title,action,description,date) VALUES (?,?,?,?,?,NOW());",
											UUID,
											smap["Type"].(string),
											smap["Title"].(string),
											smap["Action"].(string),
											smap["Description"].(string));
							responseJSON = common.JSONPackageMessage("Success","Feedback added",err!=nil);
							
						}
						w.Header().Set("Content-Type","application/json");
						w.Write(responseJSON.([]byte));
					}else{
						responseJSON = common.JSONPackageMessage("Failure","Wrong JSON formatting",err);
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

func GetRecentFeedback(w http.ResponseWriter, r *http.Request){
	// get N number of recent feedbacks
	num := mux.Vars(r)["num"];
	returnstring := "";
	rows, err := common.GetDB().Query("SELECT HEX(UUID),type,title,action,description,date FROM `"+common.GetConfig().DBTable+"` ORDER BY `Id` DESC LIMIT BY ?",num);
	
	
	if(err != nil){
        common.AsyncPrintln("[Error](GetRecentFeedback)");
		returnstring = string(common.JSONPackageMessage("","An error occured",false));
    }else{
		first := true;
		returnstring = returnstring+"{";
		for rows.Next() {
			if(!first){
				returnstring = returnstring+",";
			}else{
				first = false;
			}
			
			// get from row
			var UUID,Type,Title,Action,Description,Date string;
			err = rows.Scan(&UUID,&Type,&Title,&Action,&Description,&Date);
			if(err != nil){ // check for errors
				common.AsyncPrintln("[Error](GetRecentFeedback)");
			}else{
				// package json
				returnstring = returnstring+"{\"UUID\":\""+UUID+"\",\"Type\":\""+Type+"\",\"Title\":\""+Title+"\",\"Action\":\""+Action+"\",\"Description\":\""+Description+"\",\"Date\":\""+Date+"\"}"
			}
		}
		
		returnstring = returnstring+"}";
	}
	
	// return data
	w.Header().Set("Content-Type","application/json");
	w.Write([]byte(returnstring));
}