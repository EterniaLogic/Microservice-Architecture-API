package Feeds

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

import "../../common"
import "net/http"
import "database/sql"

func Start(scanner *common.CommandScanner){
	common.LogToFile("FeedsLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Video Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	
	//common.ListenToRegister(RegisterListener); // listen to user logins
	InitNATListener();
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}

func InitNATListener(){
// TODO: remove if migrating to single-DB
	common.GetNATSClientJSON().Subscribe("Video.NewVideo", func(subj string, reply string, msg *common.VideoData) {
		go func(){
			// add video to DB
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"_videos` (`VID`,`UUID`,`username`,`description`,`longitude`,`latitude`,`date`,`location`) VALUES (UNHEX(?),UNHEX(?),?,?,?,?,NOW(),?)",
							msg.VID,msg.UUID,msg.Username,msg.Description,msg.Longitude,msg.Latitude,msg.Location);
		}();
	});
}


func GetRecentVideos(min int,max int,w http.ResponseWriter, r *http.Request){
	// Get recent videos from the DB
	// max-min, # of followed users
	// most recent videos, only sort by followed UUIDs
	// 	Then sort them based on most recent date
	
	if test,UUID := common.RESTFulTestGetToken(w,r); test {
		//  `UUID`=UNHEX("c26720cab76611e5972c04018e7c6601") OR
		outputs, err := common.GetRowColumn(common.GetConfig().DBTable, UUID, "HEX(`followed_UUIDs`)","UUID");
		returnstring := "";
		if(!err){
			// modify the row (if it does need to be modified)
			ORIDs := "";
			numfollow := len(outputs)/32;
			if(numfollow > 0){
				for i:=0; i<len(outputs); i=i+32 {
					strid := outputs[i:i+32];
					if(i>0){
						ORIDs = ORIDs + " OR ";
					}
					ORIDs = ORIDs + "`UUID`=UNHEX('"+strid+"')";
				}
				
				// execute query and return videos
				rows, err := common.GetDB().Query("SELECT HEX(VID),username,description,date FROM `"+common.GetConfig().DBTable+"_videos` WHERE "+ORIDs+" ORDER BY `date` DESC LIMIT BY ?,?;",min,max);			
				if(err == nil){
					returnstring = GenJSONVidRows(rows);
				}else{
					returnstring = "{\"Type\":\"Failure\",\"Message\":\"An error occured!\"}";
				}
			}else{
				returnstring = "{\"Type\":\"Failure\",\"Message\":\"You are not following anybody\"}";
			}				
		}else{
			returnstring = "{\"Type\":\"Failure\",\"Message\":\"You are not following anybody\"}";
		}
		
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte(returnstring));
	}
}


func GenJSONVidRows(rows *sql.Rows)(returnstring string){
	first := true;
	returnstring = returnstring+"{\"Videos\":[";
	for rows.Next() {
		if(!first){
			returnstring = returnstring+",";
		}else{
			first = false;
		}
		
		// get from row
		var VID,Username,Description,Views,Date string;
		err := rows.Scan(&VID,&Username,&Description,&Views,&Date);
		if(err != nil){ // check for errors
			common.AsyncPrintln("[Error](GetRecentVideos)");
		}else{
			// package json
			returnstring = returnstring+"{\"VID\":\""+VID+"\",\"Username\":\""+Username+"\",\"Description\":\""+Description+"\",\"Description\":\""+Description+"\",\"Views\":\""+Views+"\",\"Date\":\""+Date+"\"}"
		}
	}
	
	returnstring = returnstring+"]}";
	return;
}