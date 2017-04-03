package Comment

// User: Brent Clancy (EterniaLogic)
// Date: 1/25/2015

import "../../common"
import "database/sql"
//import "strconv"
import "net/http"
import "encoding/json"

func Start(scanner *common.CommandScanner){
	common.LogToFile("CommentLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Comment Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}

func PostComment(w http.ResponseWriter, r *http.Request, VID int64, Comment string){
	ContentType := r.Header["Content-Type"];
	
	// get the Token from the header
	if(len(ContentType) > 0){
		if(common.RESTFulTestToken(r)){
			if(common.RESTFulGetToken(r) != ""){
				UUID := common.GetIDFromToken(common.RESTFulGetToken(r));
				if(ContentType[0] == "application/json" && UUID != ""){
					common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"` (`VID`,`UUID`,`comment`,`date`) VALUES (?,UNHEX(?),?,NOW());",
							VID,UUID,Comment);
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


func GenJSONRows(rows *sql.Rows)(returnstring string){
	first := true;
	returnstring = returnstring+"{\"Comments\":[";
	if(rows != nil){
		for rows.Next() {
			if(!first){
				returnstring = returnstring+",";
			}else{
				first = false;
			}
			
			// get from row
			var UUID,comment,Date string;
			err := rows.Scan(&UUID,&comment,&Date);
			if(err != nil){ // check for errors
				common.AsyncPrintln("[Error](GenJSONRows) sql scan column error");
			}else{
				// package json
				commentX := common.CommentX{UUID:UUID, Comment:comment, Date:Date};
				jsonx,_ := json.Marshal(commentX);
				returnstring = returnstring + string(jsonx);
			}
		}
	}
	
	returnstring = returnstring+"]}";
	return;
}