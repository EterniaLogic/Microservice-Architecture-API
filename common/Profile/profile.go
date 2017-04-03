package Profile

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

import "../../common"

func Start(scanner *common.CommandScanner){
	common.LogToFile("ProfileLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Profile Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	
	// common.ListenToLogin(Profile.LoginListener); // listen to user logins
	common.ListenToRegister(RegisterListener); // listen to user logins
	VideoWatchListener();
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}

func RegisterListener(ID string){
	CreateBlankProfileIfNotExists(ID);
}

func VideoWatchListener(){
// TODO: remove if migrating to single-DB
	common.GetNATSClientJSON().Subscribe("Video.Watch", func(subj string, reply string, msg *common.VideoView) {
		go func(){
			// add video to DB
			common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"_viewedvideos` (`VID`,`UUID`) VALUES (UNHEX(?),UNHEX(?))",
							msg.VID,msg.UUID);
		}();
	});
}

// Check if a profile exists, then create it if it does not exist.
// return true if it did not exist
func CreateBlankProfileIfNotExists(ID string)(bool){
	if(!CheckProfileExists(ID)){
		//picture, _ := ioutil.ReadFile("default.gif");
		picture := "";
		common.GetDB().Exec("INSERT INTO `"+common.GetConfig().DBTable+"` (UUID,picture) VALUES (UNHEX(?),?)",ID,picture);
		return true
	}
	
	return false
}

// check if a profile exists
func CheckProfileExists(ID string)(bool){
	_,err := common.GetRowColumn(common.GetConfig().DBTable,ID, "UUID","UUID");
	return !err;
}