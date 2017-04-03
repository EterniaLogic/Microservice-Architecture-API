package common

// Defines refactored functions to speed up Gorilla mux 
//	and RESTFul server operations.

import (
	//"log"
    "net/http"
	"strconv"
	"io/ioutil"
	"github.com/gorilla/mux"
)


func RESTFulTestToken(r *http.Request)(bool){
	Token := r.Header["Gstx"];
	
	if(len(Token) > 0){
		return true;
	}else{
		return false;
	}
}


// Output all header values
func PrintHeadersValues(r *http.Request){
	for key,value := range r.Header {
		AsyncPrintln(key);
		for x,y := range value {
			AsyncPrintln("   "+strconv.Itoa(x)+"    "+string(y));
		}
	}
}


func RESTFulGetToken(r *http.Request)(string){
	Token := r.Header["Gstx"];
	if(len(Token) > 0){
		return Token[0];
	}else{
		return "";
	}
}


func RESTFulTestGetTokenCType(w http.ResponseWriter, r *http.Request)(bool,string){
	ContentType := r.Header["Content-Type"];
	
	// get the Token from the header
	if(len(ContentType) > 0){
		if(RESTFulTestToken(r)){
			if(RESTFulGetToken(r) != ""){
				UUID := GetIDFromToken(RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && UUID != ""){
				
				
				return true,UUID;
				
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
	
	return false,"";
}

func RESTFulTestGetToken(w http.ResponseWriter, r *http.Request)(bool,string){	
	// get the Token from the header
	if(RESTFulTestToken(r)){
		if(RESTFulGetToken(r) != ""){
			UUID := GetIDFromToken(RESTFulGetToken(r));
			if(UUID != ""){
				
			return true,UUID;
			
			}else{
				w.Header().Set("Content-Type","application/json");
				w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Unknown token 2\"}"));
			}
		}else{
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Unknown Token\"}"));
		}
	}else{
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"No token provided\"}"));
	}
	
	return false,"";
}


func RESTFulNotFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Write([]byte("{\"error\":\"Not found\",\"Status\":404}"));
}


func RESTFulGet(w http.ResponseWriter, r *http.Request, table string, column string, jsontag string,idcol string){
	RESTFulGetMux(w,r,"uid",table,column,jsontag,idcol);
}


// func(w http.ResponseWriter, r *http.Request){common.RESTFulPut(w, r, "profile", "email", "Email")}
func RESTFulPut(w http.ResponseWriter, r *http.Request, table string, column string, jsonindex string){
	RESTFulPutMux(w,r,"","UUID",table,column,jsonindex);
}


func RESTFulGetMux(w http.ResponseWriter, r *http.Request, muxr string, table string, column string, jsontag string, idcol string){
	ID := mux.Vars(r)[muxr];
	id:=DecryptID(ID);
	responseJSON := JSONGetIDColumn(table, id, column, jsontag, idcol);
	w.Header().Set("Content-Type","application/json");
	w.Write(responseJSON);
}


func RESTFulGetPrivateMux(w http.ResponseWriter, r *http.Request, table string, column string, jsontag string,idcol string){
	ContentType := r.Header["Content-Type"];
	AsyncPrintln("PrivateGet: "+table+" "+column+" "+jsontag);
	if(len(ContentType) > 0){
		if(RESTFulTestToken(r)){
			AsyncPrintln("PrivateGetXToken: "+RESTFulGetToken(r));
			if(RESTFulGetToken(r) != ""){
				ID := GetIDFromToken(RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && ID != ""){
					responseJSON := JSONGetIDColumn(table, ID, column, jsontag,idcol);
					w.Header().Set("Content-Type","application/json");
					w.Write(responseJSON);
				}else{
					w.Header().Set("Content-Type","application/json");
					w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Bad header Token/MIME\"}"));
				}
			}else{
				w.Header().Set("Content-Type","application/json");
				w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Bad header Token\"}"));
			}
		}else{
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"No token provided\"}"));
		}
	}else{
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"No Token provided\"}"));
	}
}


// func(w http.ResponseWriter, r *http.Request){common.RESTFulPut(w, r, "profile", "email", "Email")}
func RESTFulPutMux(w http.ResponseWriter, r *http.Request, muxr string, byid string, table string, column string, jsonindex string){
	ContentType := r.Header["Content-Type"];
	
	if(len(ContentType) > 0){
		if(RESTFulTestToken(r)){
			if(RESTFulGetToken(r) != ""){
				ID := GetIDFromToken(RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && ID != ""){
					data,_ := ioutil.ReadAll(r.Body);
					tg,err := JSONtoMap(string(data));
					txt,err2 := GetJSONMapString(tg,jsonindex);
					
					if(!err && !err2){
						strA := "";
						strB := "";
						if(byid == "UUID"){
							strA = "UNHEX(";
							strB = ")";
						}else{
							ID = mux.Vars(r)[muxr];
						}
						GetDB().Exec("UPDATE `"+table+"` SET `"+column+"`=? WHERE `"+byid+"`="+strA+"?"+strB+";",txt,ID);
						
						q := "UPDATE `"+table+"` SET `"+column+"`='"+txt+"' WHERE `"+byid+"`="+strA+ID+strB+";";
						DBAsyncPrintln("(RESTFulPutMux): "+q);
					
						w.Header().Set("Content-Type","application/json");
						w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Set "+jsonindex+" to "+txt+"\"}"));
					}else{
						AsyncPrintln("RESTFulPutMux JSON ERROR:     "+string(data)+"   "+byid+" "+table+" "+column+" "+jsonindex);
						
						w.Header().Set("Content-Type","application/json");
						w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"JSON Error\"}"));
					}
				}else{
					w.Header().Set("Content-Type","application/json");
					w.Write([]byte("{\"Type\":\"Failure\", \"Message\":\"Bad header Token/MIME\"}"));
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


func RESTFulGetPutMux(router *mux.Router, muxr string, byIDcolumn string, prefix string, route string, dbcolumn string, JSONkey string, idcol string){
	router.HandleFunc(prefix+route+"/{"+muxr+"}", func(w http.ResponseWriter, r *http.Request){RESTFulGetMux(w, r, muxr, GetConfig().DBTable, dbcolumn, JSONkey, idcol)}).Methods("GET");
	router.HandleFunc(prefix+route, func(w http.ResponseWriter, r *http.Request){RESTFulPutMux(w, r, "",byIDcolumn, GetConfig().DBTable, dbcolumn, JSONkey)}).Methods("PUT");
}

func RESTFulGetPut2Mux(router *mux.Router, table string, muxr string, byIDcolumn string, prefix string, route string, dbcolumn string, JSONkey string, idcol string){
	router.HandleFunc(prefix+route+"/{"+muxr+"}", func(w http.ResponseWriter, r *http.Request){RESTFulGetMux(w, r, muxr, table, dbcolumn, JSONkey, idcol)}).Methods("GET");
	router.HandleFunc(prefix+route+"/{"+muxr+"}", func(w http.ResponseWriter, r *http.Request){RESTFulPutMux(w, r, muxr, byIDcolumn, table, dbcolumn, JSONkey)}).Methods("PUT");
}


// common.RESTFulGetPutPrivate(router, prefix, "/email", "email", "Email");
func RESTFulGetPutPrivate(router *mux.Router, prefix string, route string, dbcolumn string, JSONkey string, idcol string){
	router.HandleFunc(prefix+route, func(w http.ResponseWriter, r *http.Request){RESTFulGetPrivateMux(w, r, GetConfig().DBTable, dbcolumn, JSONkey, idcol)}).Methods("GET");
	router.HandleFunc(prefix+route, func(w http.ResponseWriter, r *http.Request){RESTFulPut(w, r, GetConfig().DBTable, dbcolumn, JSONkey)}).Methods("PUT");
}


func RESTFulGetPut(router *mux.Router, prefix string, route string, dbcolumn string, JSONkey string, idcol string){
	router.HandleFunc(prefix+route+"/{uid}", func(w http.ResponseWriter, r *http.Request){RESTFulGet(w, r, GetConfig().DBTable, dbcolumn, JSONkey, idcol)}).Methods("GET");
	router.HandleFunc(prefix+route, func(w http.ResponseWriter, r *http.Request){RESTFulPut(w, r, GetConfig().DBTable, dbcolumn, JSONkey)}).Methods("PUT");
}