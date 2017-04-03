package Administration

import (
    "net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"../../common"
)

type AdminREST struct{
	router *mux.Router;
	prefix string;
}

func NewAdminREST()(t AdminREST){
	// initialize the server data
	t = AdminREST{};
	version := "v1";
	t.prefix = "/api/"+version+"/admin";
	t.router = mux.NewRouter().StrictSlash(true);
	
	//t.AddPutRoute("/user/{uid}",); // modify user
	//t.AddPutRoute("/user/level/{uid}",); // modify user level
	//t.AddGetRoute("/user/{uid}",); // get user info
	//t.AddDelRoute("/user/{uid}",); // remove user
	
	t.AddPostRoute("", RESTAdminLogin); // login
	
	return;
}

func (t AdminREST) AddPutRoute(loc string, tfunc func(http.ResponseWriter, *http.Request)){
	t.router.HandleFunc(t.prefix+loc, tfunc).Methods("PUT");
}

func (t AdminREST) AddPostRoute(loc string, tfunc func(http.ResponseWriter, *http.Request)){
	t.router.HandleFunc(t.prefix+loc, tfunc).Methods("POST");
}

func (t AdminREST) AddGetRoute(loc string, tfunc func(http.ResponseWriter, *http.Request)){
	t.router.HandleFunc(t.prefix+loc, tfunc).Methods("GET");
}

func (t AdminREST) AddDelRoute(loc string, tfunc func(http.ResponseWriter, *http.Request)){
	t.router.HandleFunc(t.prefix+loc, tfunc).Methods("DELETE");
}

func (t AdminREST) Start(port string){
	// start http server
	err := http.ListenAndServe(":"+port, t.router);
	if(err != nil){
		panic(err);
	}
}

func RESTAdminLogin(w http.ResponseWriter, r *http.Request){
	body,_:=ioutil.ReadAll(r.Body);
	smap,err := common.JSONtoMap(string(body));
	
	// "Username", "Password"
	if(!err && admins[r.RemoteAddr].Username == "" && smap["Username"] != nil){
		admins[r.RemoteAddr] = Admin{Username: smap["Username"].(string)};
		if(admins[r.RemoteAddr].UserLogin(smap["Username"].(string), smap["Password"].(string))){
			// logged in
			w.Header().Set("Content-Type","application/json");
			w.Write(common.JSONPackageMessage("Logged in", "", true));
		}else{
			w.Header().Set("Content-Type","application/json");
			w.Write(common.JSONPackageMessage("", "Username/Password incorrect", false));
		}
	}else{
		// is this person already logged in?
		common.AsyncPrintln("[JSON Error] RESTAdminLogin");
	}
}