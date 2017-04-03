package Feeds 

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "net/http"
	"github.com/gorilla/mux"
	"../../common"
	"strings"
	"strconv"
)


// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/feeds";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
		
	// Get/Modify
	router.HandleFunc(prefix+"/follow/{uid}", HandleFollowUser).Methods("POST"); // follow user
	router.HandleFunc(prefix+"/follow/{uid}", HandleUnfollowUser).Methods("DELETE"); // unfollow user
	router.HandleFunc(prefix+"/follow", HandleGetFollowingUsers).Methods("GET"); // Get followed users
	router.HandleFunc(prefix+"/recent/{min}/{max}", HandleRecentFeeds).Methods("GET");
		
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}


func HandleFollowUser(w http.ResponseWriter, r *http.Request){
	uid := mux.Vars(r)["uid"];
	
	// Get user token from the header
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		// follow that user (if it hasnt been followed already)
		if(common.RowExists(UUID, "UUID", "",true)){
			outputs, err := common.GetRowColumn(common.GetConfig().DBTable, UUID, "HEX(`followed_UUIDs`)","UUID");
			if(!err){
				// modify the row (if it does need to be modified)
				instr := false; // is the userid in the string?
				for i:=0; i<len(outputs); i=i+32 {
					str := outputs[i:i+32];
					if(strings.ToLower(str) == strings.ToLower(uid)){
						instr = true;
					}
				}
				
				// Add to the index
				if(!instr){
					outputs += uid;
					common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"` SET `followed_UUIDs`=UNHEX(?) WHERE `UUID`=UNHEX(?);",outputs,UUID);
					w.Header().Set("Content-Type","application/json");
					w.Write([]byte("{\"Type\":\"Success\",\"Message\":\"Followed user\""));
				}
			}
		}else{
			// add a new row with the ID
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"` (`UUID`,`followed_UUIDs`) VALUES (UNHEX(?),UNHEX(?));",UUID,uid);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Success\",\"Message\":\"Followed user\""));
		}
	}
}


// Follow the specified uid, your uid is from the token
func HandleUnfollowUser(w http.ResponseWriter, r *http.Request){
	uid := mux.Vars(r)["uid"];
	
	// Get user token from the header
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		// follow that user (if it hasnt been followed already)
		outputs, err := common.GetRowColumn(common.GetConfig().DBTable, UUID, "HEX(`followed_UUIDs`)","UUID");
		newoutputs := "";
		if(!err){
			// modify the row (if it does need to be modified)
			instr := false; // is the userid in the string?
			for i:=0; i<len(outputs); i=i+32 {
				str := outputs[i:i+32];
				if(strings.ToLower(str) == strings.ToLower(uid)){
					instr = true;
				}else{
					newoutputs = newoutputs+str;
				}
			}
			
			// remove to the index
			if(instr){
				common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"` SET `followed_UUIDs`=UNHEX(?) WHERE `UUID`=UNHEX(?);",newoutputs,UUID);
				w.Header().Set("Content-Type","application/json");
				w.Write([]byte("{\"Type\":\"Success\",\"Message\":\"Unfollowed user\""));
			}
		}
	}
}

func HandleGetFollowingUsers(w http.ResponseWriter, r *http.Request){
	// Get user token from the header
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		outputs, err := common.GetRowColumn(common.GetConfig().DBTable, UUID, "HEX(`followed_UUIDs`)","UUID");
		
		if(!err){
			// modify the row (if it does need to be modified)
			numfollow := len(outputs)/32;
			jsonoutput := "{";
			jsonoutput = jsonoutput + "\"Number\":"+strconv.Itoa(numfollow)+",";
			jsonoutput = jsonoutput + "\"Following\":[";
			for i:=0; i<len(outputs); i=i+32 {
				str := outputs[i:i+32];
				// loop through and produce JSON
				if(i>0) {
					jsonoutput = jsonoutput+",";
				}
				jsonoutput = jsonoutput + "\""+str+"\"";
			}
			jsonoutput = jsonoutput + "]}";
			
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte(jsonoutput));
		}else{
			common.DBAsyncPrintln("[HandleGetFollowingUsers] SQL Error: "+outputs);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Failure\",\"Message\":\"SQL Error\"}"));
		}
	}
}

func HandleRecentFeeds(w http.ResponseWriter, r *http.Request){
	mins := mux.Vars(r)["min"];
	maxs := mux.Vars(r)["max"];
	
	min := 0;
	max := 50;
	
	mini,err1 := strconv.Atoi(mins);
	maxi,err2 := strconv.Atoi(maxs);
	
	// finish min/max comparison
	if(err1 == nil && err2 == nil){
		min,max = common.MinMaxMixer(mini,maxi);
		
		
		// Get user token from the header
		GetRecentVideos(min,max,w,r);
	}
}
