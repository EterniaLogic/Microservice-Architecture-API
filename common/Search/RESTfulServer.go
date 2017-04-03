package Search 

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
	prefix := "/api/"+version+"/search";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
		
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
	
	// Get/Modify
	router.HandleFunc(prefix, HandleSearch).Methods("POST");
	
	
	//RESTFulGetPut2Mux(router *mux.Router, table, mux string, byIDcolumn string, prefix string, route string, dbcolumn string, JSONkey string, idcol string){
	common.RESTFulGetPut2Mux(router,common.GetConfig().DBTable+"_videos", "vid", "VID", prefix, "/tags", "tags", "Tags","VID");
	
	//common.RESTFulGetPut2Mux(router, "vid", "VID", prefix, "/description", "description", "Description", "VID");
	
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}

func HandleSearch(w http.ResponseWriter, r *http.Request){
	// Handle Custom JSON for a search
	returnstring := "";
	body,_:=ioutil.ReadAll(r.Body);
	smap,err := common.JSONtoMap(string(body));
	
	if(err){
		// JSON error
		//common.AsyncPrintln("[Error](HandleSearch)");
		returnstring = string(common.JSONPackageMessage("","No videos available",false));
		
	}else{
		// JSON transfered over, inspect the output
		search := smap["Search"].(string);
		
		returnstring = PerformSearch(search);
	}
	
	w.Header().Set("Content-Type","application/json");
	w.Write([]byte(returnstring));
}