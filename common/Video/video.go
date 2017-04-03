package Video

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

import "../../common"
import "encoding/json"
import "database/sql"
import "strconv"
import "net/http"

func Start(scanner *common.CommandScanner){
	common.LogToFile("VideoLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Video Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	//common.ListenToRegister(RegisterListener); // listen to user logins
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}

// user likes a video, if link clicked again, unlike video
func LikeVideo(w http.ResponseWriter, VID string, UUID string){
	var IDx int64;
	vid,_ := strconv.ParseInt(VID,10,64);
	common.GetDB().QueryRow("SELECT vid FROM `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `UUID`=UNHEX(?) AND `islike`=0 LIMIT 1;",vid,UUID).Scan(&IDx);
	
	// if islike, set to like
	if(IDx == vid){
		common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"_liked` SET `islike`=0 WHERE `vid`=? AND `UUID`=UNHEX(?) LIMIT 1;",vid,UUID);
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Liked Video\"}"));
	}else{
		// else, add like
		if(!common.RowExists2(VID,"vid",UUID,"UUID","_liked")){
			// like this video
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"_liked` (`vid`,`UUID`,`islike`) VALUES (?,UNHEX(?),1)",vid,UUID);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Liked Video\"}"));
		}else{
			// unlike this video (delete row)
			common.GetDB().Exec("DELETE FROM `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `UUID`=UNHEX(?)",vid,UUID);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Unliked Video\"}"));
		}
	}
}

// user dislikes a video, if link clicked again, un-dislike video
func DislikeVideo(w http.ResponseWriter, VID string, UUID string){
	var IDx int64;
	vid,_ := strconv.ParseInt(VID,10,64);
	common.GetDB().QueryRow("SELECT vid FROM `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `UUID`=UNHEX(?) AND `islike`=1 LIMIT 1;",vid,UUID).Scan(&IDx);
	
	// if islike, set to dislike
	if(IDx == vid){
		common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"_liked` SET `islike`=0 WHERE `vid`=? AND `UUID`=UNHEX(?) LIMIT 1;",vid,UUID);
		w.Header().Set("Content-Type","application/json");
		w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Disliked Video\"}"));
	}else{
		// else, add dislike
		if(!common.RowExists2(VID,"vid",UUID,"UUID","_liked")){
			// dislike this video
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"_liked` (`vid`,`UUID`,`islike`) VALUES (?,UNHEX(?),0)",vid,UUID);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Disliked Video\"}"));
		}else{
			// undislike this video (delete row)
			common.GetDB().Exec("DELETE FROM `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `UUID`=UNHEX(?)",vid,UUID);
			w.Header().Set("Content-Type","application/json");
			w.Write([]byte("{\"Type\":\"Success\", \"Message\":\"Undisliked Video\"}"));
		}
	}
}

// user dislikes a video, if link clicked again, un-dislike video
func PostVideo(UUID string, VidLink string, username string, description string, longitude string, latitude string, location string)(msg string, err bool){
	// determine if this video exists or not
	// INSERT INTO `video` (`VID`,`VidLink`,`UUID`,`username`,`description`,`longitude`,`latitude`,`date`,`location`) VALUES (UNHEX("EDnAf2w2v-Y"),"https://www.youtube.com/watch?v=EDnAf2w2v-Y",UNHEX("13a6dcf3d52311e58aee0401a8d8aa01"),"EterniaLogic","Our new offsite storage server will hold 160TB (expandable to 264TB)!","33º34'57.32\" N","101º52'42.96\" W",NOW(),"Lubbock, TX");
	
	if(!common.RowExists(VidLink,"VidLink","",false)){
		cque,errx := common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"` (`VidLink`,`UUID`,`username`,`description`,`longitude`,`latitude`,`date`,`location`) VALUES (?,UNHEX(?),?,?,?,?,NOW(),?);",
							VidLink,UUID,username,description,longitude,latitude,location);
		
		if(errx != nil){
			common.DBAsyncPrintln("PostVideo: {"+VidLink+",   "+UUID+",   "+username+",   "+description+",   "+longitude+",   "+latitude+",   "+location+"} ERROR="+errx.Error());
			err = true;
			msg = "Error storing the video!";
		}else{
			// Replicate data for Search, Feeds
			// TODO: remove if migrating to single-DB
			idp, _ := cque.LastInsertId();
			Vdata := common.VideoData{};
			Vdata.VID = idp;
			Vdata.VidLink = VidLink;
			Vdata.UUID = UUID;
			Vdata.Username = username;
			Vdata.Description = description;
			Vdata.Longitude = longitude;
			Vdata.Latitude = latitude;
			Vdata.Location = location;
			common.GetNATSClientJSON().Publish("Video.NewVideo",Vdata);
			
			
			err = false;
			msg = "Video created";
		}
	}else{
		err = true;
		msg = "Video exists";
	}
	return;
}

// returns information related to the video
//	including likes, dislikes
// 	UUID - is if a user is logged in, else no like/dislike return allowed for this user
func GetVideo(VID string, UUID string, loggedin bool)(msg string){
	// Get the video from the database
	vid,_ := strconv.ParseInt(VID,10,64);
	var username, description, longitude, latitude, location, date string;
	var sponsoredx int;
	var views int;
	// &username,&description, &longitude, &latitude, &location, &sponsored, &date, &views
	errx := common.GetDB().QueryRow("SELECT username,description,longitude,latitude,location,`date`,sponsored,views FROM `"+common.GetConfig().DBTable+"` WHERE `VID`=? LIMIT 1;",VID).Scan(&username,&description, &longitude, &latitude, &location, &date, &sponsoredx, &views);
	
	sponsored := false;
	if(sponsoredx == 1){
		sponsored = true;
	}
	
	if(errx == nil){
		vdata := common.VideoData{VID:vid, UUID:UUID, Username:username, Description:description, Longitude:longitude, Latitude:latitude, Location:location, Date:date, Sponsored:sponsored, Views:views};
		if(loggedin){
			var islike int;
			common.GetDB().QueryRow("SELECT `islike` FROM `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `UUID`=UNHEX(?) LIMIT 1;",vid,UUID).Scan(&islike);
			
			// if islike, set to dislike
			if(islike==1){
				vdata.Like = true;
			}else{
				vdata.Dislike = true;
			}
			
			// send out that this video has been viewed
			// Replicate data for Search, Feeds
			// TODO: remove if migrating to single-DB
			Vdata := common.VideoView{VID: vid, UUID: UUID};
			common.GetNATSClientJSON().Publish("Video.Watch",Vdata);
		}
		
		// count likes for the video
		common.GetDB().QueryRow("SELECT COUNT(`VID`) from `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `islike`=1;",VID).Scan(&vdata.Likes);
		common.GetDB().QueryRow("SELECT COUNT(`VID`) from `"+common.GetConfig().DBTable+"_liked` WHERE `vid`=? AND `islike`=0;",VID).Scan(&vdata.Dislikes);
		
		tmp, err2 := json.Marshal(vdata);
		if err2 != nil {
			common.AsyncPrintln("JSON error (GetVideo) "+err2.Error());
			msg = "{\"Type\":\"Failure\",\"Message\":\"error occured!\"}";
		}else{
			msg=string(tmp);
		}
		
	}else{
		common.DBAsyncPrintln("SQL error (GetVideo) "+errx.Error());
		msg = "{\"Type\":\"Failure\",\"Message\":\"Video does not exist\"}";
	}
	
	return;
}

func GetTopVideos(min int, max int)(returnstring string){
	// make sure input is correct:
	min,max = common.MinMaxMixer(min,max);
	
	rows, err := common.GetDB().Query("SELECT VID,username,description,views,`date` FROM `"+common.GetConfig().DBTable+"` ORDER BY `views` DESC LIMIT ?,?;",min,max);
	
	
	if(err != nil){
        common.AsyncPrintln("[Error](GetRecentVideos)");
		returnstring = string(common.JSONPackageMessage("","An error occured",false));
    }else{
		returnstring = GenJSONRows(rows);
	}
	
	return;
}

func GetRecentVideos(min int, max int)(returnstring string){
	// make sure input is correct:
	min,max = common.MinMaxMixer(min,max);
	
	rows, err := common.GetDB().Query("SELECT VID,username,description,views,`date` FROM `"+common.GetConfig().DBTable+"` ORDER BY `date` DESC LIMIT ?,?;",min,max);
	
	
	if(err != nil){
        common.AsyncPrintln("[Error](GetRecentVideos)");
		returnstring = string(common.JSONPackageMessage("","An error occured",false));
    }else{
		returnstring = GenJSONRows(rows);
	}
	
	return;
}


func GenJSONRows(rows *sql.Rows)(returnstring string){
	first := true;
	returnstring = returnstring+"{\"Videos\":[";
	for rows.Next() {
		if(!first){
			returnstring = returnstring+",";
		}else{
			first = false;
		}
		
		// get from row
		var VID int;
		var Username,Description,Views,Date string;
		err := rows.Scan(&VID,&Username,&Description,&Views,&Date);
		if(err != nil){ // check for errors
			common.AsyncPrintln("[Error](GenJSONRows) "+err.Error());
		}else{
			// package json
			vid := strconv.Itoa(VID);
			returnstring = returnstring+"{\"VID\":\""+vid+"\",\"Username\":\""+Username+"\",\"Description\":\""+Description+"\",\"Description\":\""+Description+"\",\"Views\":\""+Views+"\",\"Date\":\""+Date+"\"}"
		}
	}
	
	returnstring = returnstring+"]}";
	return;
}

