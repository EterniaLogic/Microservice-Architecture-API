package Profile 

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "net/http"
	"github.com/gorilla/mux"
	"../../common"
)


// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/profiles";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
	
	// PUT functions (Requires a header with Token, specifically for the current user)
	//router.HandleFunc(prefix, SetUserProfile).Methods("PUT");
	
	// get data
	router.HandleFunc(prefix+"/{uid}", GetUserProfile).Methods("GET");
		
	// Get/Modify DB rows, based on DB ID (Private)
	common.RESTFulGetPutPrivate(router, prefix, "/email", "email", "Email","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/firstname", "firstname", "Firstname","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/lastname", "lastname", "Lastname","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/gender", "gender", "Gender","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/city", "city", "City","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/state", "state", "State","UUID");
	common.RESTFulGetPutPrivate(router, prefix, "/country", "country", "Country","UUID");
	
	
	// public
	common.RESTFulGetPut(router, prefix, "/picture", "picture", "Picture","UUID");
	common.RESTFulGetPut(router, prefix, "/bio", "bio", "Bio","UUID");
	common.RESTFulGetPut(router, prefix, "/facebook", "facebook", "Facebook","UUID");
	common.RESTFulGetPut(router, prefix, "/twitter", "twitter", "Twitter","UUID");
	common.RESTFulGetPut(router, prefix, "/skype", "skype", "Skype","UUID");
	common.RESTFulGetPut(router, prefix, "/whatsapp", "whatsapp", "Whatsapp","UUID");
	common.RESTFulGetPut(router, prefix, "/snapchat", "snapchat", "Snapchat","UUID");
	common.RESTFulGetPut(router, prefix, "/instagram", "instagram", "Instagram","UUID");
	common.RESTFulGetPut(router, prefix, "/kik", "kik", "Kik","UUID");
	common.RESTFulGetPut(router, prefix, "/website", "website", "Website","UUID");
	common.RESTFulGetPut(router, prefix, "/gear", "gear", "Gear","UUID");
	common.RESTFulGetPut(router, prefix, "/birthday", "birthday", "Birthday","UUID");
		
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}

// Handle get Profile
// GET /api/v1/profile/{uid}
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["uid"];
	responseJSON := JSONGetUserProfile(id);
	w.Header().Set("Content-Type","application/json");
	w.Write(responseJSON);
}