package Auth 

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
    "net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"../../common"
)

// Start up a HTTP server
func StartRESTfulServer(){
    common.AsyncPrintln("Starting RESTful Server...");
    
	version := "v1";
	prefix := "/api/"+version+"/auth";
	
	router := mux.NewRouter().StrictSlash(true);
	// router.Host("site.com"); // prevent Cross-Site?
	
	// 404 Page not found
	router.NotFoundHandler = http.HandlerFunc(common.RESTFulNotFound);
	
	// register
	router.HandleFunc(prefix+"/user", UserCreate).Methods("POST");
	
	// get data
	router.HandleFunc(prefix+"/user/name/{uid}", func(w http.ResponseWriter, r *http.Request){common.RESTFulGet(w, r, common.GetConfig().DBTable, "username", "Username","UUID")}).Methods("GET");
	router.HandleFunc(prefix+"/user/lastlogin/{uid}", func(w http.ResponseWriter, r *http.Request){common.RESTFulGet(w, r, common.GetConfig().DBTable, "lastlogin", "LastLogin","UUID")}).Methods("GET");
	
	// Verify email
	router.HandleFunc(prefix+"/verify/{vcode}", UserVerifyEmail).Methods("GET"); // using GET for a normal browser
	
	// Append OAuth2 token
	router.HandleFunc(prefix+"/oauth", func(w http.ResponseWriter, r *http.Request){common.RESTFulPut(w, r, common.GetConfig().DBTable, "oauth2_token", "OAuth2Token")}).Methods("PUT");
	
	// login/logout
	router.HandleFunc(prefix+"/login", UserLogin).Methods("POST");
	router.HandleFunc(prefix+"/login/{uid}", UserLogout).Methods("DELETE");
	
	
	// Debugging & Profiling
	common.ProfilingRouterHandler(router);	
	
	// start http server
	err := http.ListenAndServe(":"+common.GetConfig().RESTfulPort, router);
	if(err != nil){
		panic(err);
	}
}

// create a user with JSON
// POST /api/v1/auth/user
// {"Username":"JohnDoe","UserPass":"test","Email":"johndoe@gmail.com"}
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var responseJSON interface{};
	body,_:=ioutil.ReadAll(r.Body);
	smap,err := JSONtoMap(string(body));
	
	if(!err){
		responseJSON = common.TestJSONInput(smap["Username"], "JSON Username key not found");
		responseJSON = common.TestJSONInput(smap["UserPass"], "JSON UserPass key not found");
		responseJSON = common.TestJSONInput(smap["Email"], "JSON Email key not found");	
		
		if(responseJSON == nil){
			msg,err := AuthRegister(smap["Username"].(string), smap["UserPass"].(string), smap["Email"].(string));
			responseJSON = common.JSONPackageMessage("Success",msg,err);
		}
	}else{
		responseJSON = common.JSONPackageMessage("Failure","Wrong JSON formatting",err);
	}
	w.Header().Set("Content-Type","application/json");
	w.Write(responseJSON.([]byte));
}

// Handle login user
// POST /api/v1/auth/user/login
// {"Username":"JohnDoe","UserPass":"test"}
// {"ID":"95341411595","Token":"8eff1b488da7fe3426f9ecaf8de1ba54","Website":"facebook.com"}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var responseJSON interface{};
	body,_:=ioutil.ReadAll(r.Body);
	smap,err := common.JSONtoMap(string(body));
	
	if(!err){
		if(smap["Username"] != nil){
			// handle Direct authentication
			responseJSON = common.TestJSONInput(smap["Username"], "JSON Username key not found");
			responseJSON = common.TestJSONInput(smap["UserPass"], "JSON UserPass key not found");
			
			if(responseJSON == nil){
				responseJSON = JSONPackageLogin(AuthUser(smap["Username"].(string), smap["UserPass"].(string), false));
			}
		}else{
			// handle OAuth2 authentication
			responseJSON = common.TestJSONInput(smap["Token"], "JSON OAuth2 Secret key not found");
			responseJSON = common.TestJSONInput(smap["Website"], "JSON Website key not found");
			
			if(responseJSON == nil){
				responseJSON = JSONPackageLogin(OAuthUser(smap["Token"].(string), smap["Website"].(string)));
			}
		}
	}else{
		responseJSON = common.JSONPackageMessage("Failure","Wrong JSON formatting",err);
	}
	w.Header().Set("Content-Type","application/json");
	w.Write(responseJSON.([]byte));
}


//
func UserVerifyEmail(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["vcode"];
	var UUID,auth_level string;
	common.GetDB().QueryRow("SELECT HEX(UUID),`auth_level` FROM `"+common.GetConfig().DBTable+"` WHERE `auth_token`=UNHEX(?);",code).Scan(&UUID,&auth_level);
	if(UUID != "" && common.ConvertStringToLevel(auth_level)==1){
	
		authlgen := common.GenAuthLevel("User",UUID);
		common.GetDB().Exec("UPDATE "+common.GetConfig().DBTable+" SET `auth_level`='User', `auth_level_gen`=UNHEX(?) WHERE `UUID`=UNHEX(?);",authlgen,UUID);	
		
		w.Header().Set("Content-Type","text/html");
		w.Write([]byte("<!DOCTYPE html><html><head><title>New User Registration Verification</title><script src='../js/verification.js'></head><body data-name='username' data-status='success'></body></html>"));
	}else{
		w.Header().Set("Content-Type","text/html");
		w.Write([]byte("<!doctype html><html><head><title>New User Registration Verification</title><script src='../js/verification.js'></head><body data-name='username' data-status='failure'></body></html>"));
	}
}


// Handle logout a user
// DELETE /api/v1/auth/login/{uid}
func UserLogout(w http.ResponseWriter, r *http.Request) {	
	id := mux.Vars(r)["uid"];
	responseJSON := JSONDeAuth(id);
	w.Header().Set("Content-Type","application/json");
	w.Write(responseJSON);
}
