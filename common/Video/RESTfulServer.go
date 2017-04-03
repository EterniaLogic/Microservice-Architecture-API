package Video 

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"../../common"
	"strconv"
)


// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/videos";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// Post a new video
	router.HandleFunc(prefix+"", HandlePostVideo).Methods("POST");
	
	// Get a video
	router.HandleFunc(prefix+"/{vid}", HandleGetVideo).Methods("GET");
	
	// Get N number of top videos (top views, max 100)
	router.HandleFunc(prefix+"/top/{num}/{tonum}", HandleTopVideos).Methods("GET");
	
	// Get N number of recent videos (Max 100)
	router.HandleFunc(prefix+"/recent/{num}/{tonum}", HandleRecentVideos).Methods("GET");
	
	// Post a like (if already liked, unlike)
	//	anybody logged in may do this
	router.HandleFunc(prefix+"/like/{vid}", HandleLikeVideo).Methods("POST");
	
	// Post/Delete a dislike, flips like
	//	anybody logged in may do this
	router.HandleFunc(prefix+"/dislike/{vid}", HandleDislikeVideo).Methods("POST");
	
	// Delete a video
	// 	only the video creator may do this
	//router.HandleFunc(prefix+"/delete/{vid}", HandleDeleteVideo).Methods("DELETE");
	
	// Get/Modify
	common.RESTFulGetPut2Mux(router,common.GetConfig().DBTable, "vid", "VID", prefix, "/description", "description", "Description", "VID");
		
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}

// multi-parameter post for video
func HandlePostVideo(w http.ResponseWriter, r *http.Request){
	var responseJSON interface{};
	
	// get the Token from the header
	if test,UUID := common.RESTFulTestGetTokenCType(w,r); test {
		body,_:=ioutil.ReadAll(r.Body);
		smap,err := common.JSONtoMap(string(body));
		
		// JSON to map error?
		if(!err){
			responseJSON = common.TestJSONInput(smap["VidLink"], "JSON VidLink key not found");
			responseJSON = common.TestJSONInput(smap["Username"], "JSON Username key not found");
			responseJSON = common.TestJSONInput(smap["Description"], "JSON Description key not found");
			responseJSON = common.TestJSONInput(smap["Longitude"], "JSON Longitude key not found");
			responseJSON = common.TestJSONInput(smap["Latitude"], "JSON Latitude key not found");
			responseJSON = common.TestJSONInput(smap["Location"], "JSON Latitude key not found");
			
			// no error? Post video
			if(responseJSON == nil){
				msg,err := PostVideo(UUID,
										smap["VidLink"].(string),
										smap["Username"].(string), 
										smap["Description"].(string), 
										smap["Longitude"].(string), 
										smap["Latitude"].(string),
										smap["Location"].(string));
				responseJSON = common.JSONPackageMessage("Success",msg,err);
			}
		}else{
			responseJSON = common.JSONPackageMessage("Failure","Wrong JSON formatting",err);
		}
		w.Header().Set("Content-Type","application/json");
		w.Write(responseJSON.([]byte));
	}
}

// multi-parameter post for video
func HandleGetVideo(w http.ResponseWriter, r *http.Request){
	VID := mux.Vars(r)["vid"];
	UUID := "";
	isLoggedIn := false;
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		UUID = common.GetIDFromToken(common.RESTFulGetToken(r));
		if(UUID != ""){
			isLoggedIn = true;
		}
	}
	
	// if logged in, get whether liked/disliked
	w.Header().Set("Content-Type","application/json");
	w.Write([]byte(GetVideo(VID,UUID,isLoggedIn)));
}

// Get top videos
func HandleTopVideos(w http.ResponseWriter, r *http.Request){
	min := mux.Vars(r)["num"];
	max := mux.Vars(r)["tonum"];
	
	mini,err1 := strconv.Atoi(min);
	maxi,err2 := strconv.Atoi(max);
	
	if(err1 == nil && err2 == nil){
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte(GetTopVideos(mini,maxi)));
	}
}

// Get recent videos for a user
func HandleRecentVideos(w http.ResponseWriter, r *http.Request){
	min := mux.Vars(r)["num"];
	max := mux.Vars(r)["tonum"];
	
	mini,err1 := strconv.Atoi(min);
	maxi,err2 := strconv.Atoi(max);
	
	if(err1 == nil && err2 == nil){
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte(GetRecentVideos(mini,maxi)));
	}
}

// user likes a video
func HandleLikeVideo(w http.ResponseWriter, r *http.Request){
	VID := mux.Vars(r)["vid"];
	
	// get the Token from the header
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		LikeVideo(w,VID,UUID);
	}
}

func HandleDislikeVideo(w http.ResponseWriter, r *http.Request){
	VID := mux.Vars(r)["vid"];
	
	// get the Token from the header
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		DislikeVideo(w,VID,UUID);
	}
}

func HandleDeleteVideo(w http.ResponseWriter, r *http.Request){
	
}