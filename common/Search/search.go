package Search

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

import "../../common"
import "strings"
import "container/list"

type SortStruct struct{
	VID string;
	occurences int;
}

func Start(scanner *common.CommandScanner){
	common.LogToFile("SearchLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Search Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	
	InitNATListener();
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}

func InitNATListener(){
// TODO: remove if migrating to single-DB
	common.GetNATSClientJSON().Subscribe("Video.NewVideo", func(subj string, reply string, msg *common.VideoData) {
		go func(){
			// add video to DB
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"_videos` (`VID`,`VidLink`,`UUID`,`username`,`description`,`longitude`,`latitude`,`date`,`location`) VALUES (UNHEX(?),?,UNHEX(?),?,?,?,?,NOW(),?)",
							msg.VID,msg.VidLink,msg.UUID,msg.Username,msg.Description,msg.Longitude,msg.Latitude,msg.Location);
		}();
	});
}


func PerformSearch(searchString string) (string){
	l1 := list.New();
	searchString=ReplaceNonAlphaneumeric(searchString);
	spaces := strings.Split(searchString," "); // split up search by spaces
	commas := strings.Split(searchString,","); // split up search by commas
	
	// Split up search and look up based on tags, gear, activity
	// Count occurences and search based on amount of "sameness"
	// only count by "Uniqueness"
	var VIDs map[string]int;
	VIDs = make(map[string]int);
	CountUniqueVid(spaces, &VIDs);	// count LIKE based on spaces
	CountUniqueVid(commas, &VIDs); // count LIKE based on commas
	
	// sort videos based on # of occurences
	SortMapValues(&VIDs, l1);
	
	// Combine found values into a JSON output
	return GenJSONRows(l1);
}


func ReplaceNonAlphaneumeric(input string)(string){
	strout := "";
	for _,c := range input {
		cint := int(c); // get raw ASCII value
		
		// compare ascii table
		if((48 <= cint && cint <= 57) || (65 <= cint && cint <= 90) || (97 <= cint && cint <= 122)){
			// add to a temp string
			strout += string(c);
		}
	}
	return strout;
}


// count unique terms based on description
func CountUniqueVid(terms []string, VIDs *map[string]int){
	for _,term := range terms {
		sterm := "%"+term+"%";
		rows, err := common.GetDB().Query("SELECT HEX(VID) FROM `"+common.GetConfig().DBTable+"_videos` WHERE `description` LIKE ? OR `tags` LIKE ? LIMIT 100;",sterm,sterm);
		common.AsyncPrintln("DEBUG: "+term);
		// sql err
		if(err == nil){
			for rows.Next() {
				// get from row
				var VID string;
				err := rows.Scan(&VID);
				common.AsyncPrintln("DEBUG 1: "+VID);
				if(err != nil){ // check for errors
					common.DBAsyncPrintln("[Error](GetRecentVideos 1) "+err.Error());
				}else{
					// add to map
					common.AsyncPrintln("DEBUG 2: "+VID);
					(*VIDs)[VID] = (*VIDs)[VID]+1;
				}
			}
		}else{
			common.DBAsyncPrintln("[Error](GetRecentVideos 2) "+err.Error());
		}
	}
}


// Sort based on # of occurences
func SortMapValues(VIDs *map[string]int, list1 *list.List){
	// Selection sort O(N^2) will be used here, but if need be a quicksort O(nlogn) can be used later.
	// Selection sort will run up to 400 items here, which is about 400^2 operations (800k ops)
	// 2GHz / 800k = 400 uS best case, if CPU only uses 1CPI
	
	redomaxcalc := true;
	max := 0;
	for key,term := range *VIDs {
		// find the max values
		
		if(redomaxcalc){ // prevent re-finding max until change
			max = 0;
			for _,term := range *VIDs {
				if(max < term){
					max = term;
					redomaxcalc = false; // wait until something changes
				}
			}
		}
		
		if(term >= max){
			//sortchannel <- key; // send data down a channel (a queue)
			list1.PushFront(key);
			delete(*VIDs, key); // remove this item from the map
			redomaxcalc = true;
		}
	}
}


// Generate JSON for output
func GenJSONRows(list1 *list.List)(returnstring string){
	first := true;
	//common.AsyncPrintln("DEBUG: (GenJSONRows) ");
	returnstring = returnstring+"{\"Videos\":[";
	for e := list1.Front(); e != nil; e = e.Next() {
		//common.AsyncPrintln("DEBUG: (GenJSONRows 2): "+e.Value.(string));
		if(!first){
			returnstring = returnstring+",";
		}else{
			first = false;
		}
		
		// get from row
		returnstring = returnstring+e.Value.(string);
	}
	
	returnstring = returnstring+"]}";
	return;
}