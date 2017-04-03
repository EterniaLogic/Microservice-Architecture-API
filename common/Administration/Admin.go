package Administration

// User: Brent Clancy (EterniaLogic)
// Date: 1/10/2016

import "../../common"
import "math"
import "time"

var admins map[string]Admin;

type Admin struct{
	UUID string;		// UUID HEX for the admin
	Username string;
	IP string;			// IP address for the admin
	IDToken string;
	IsAdmin string;		// Verified admin
	UserLevel string;	// plaintext userlevel
	UserLevelGen string;// Generated HEX userlevel
};

// Initialize the admin RESTFul server
func Start(scanner *common.CommandScanner){
	admins = map[string]Admin{};
	common.LogToFile("AdminLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Administration Server...");
	common.SetConfig("conf.json");
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	
	// onclose function for Ctrl-C or SIGTERM
	common.OnClose();
	
	restful := NewAdminREST();
	restful.Start(common.GetConfig().RESTfulPort);
}

// verify a user's level on the communication channel
func (a Admin) VerifyUser(ulevel string) bool{
	var isadmin bool;
	var tolvl bool;
	
	level := common.ConvertStringToLevel(a.UserLevel);
	wantedlevel := common.ConvertStringToLevel(ulevel);
	
	if(ulevel=="CommentModerator" && math.Mod(float64(level),2)!=0){ // CommentMod
		tolvl = true;
	}else if(ulevel=="VideoModerator" && math.Mod(float64(level),3)!=0){
		tolvl = true;
	}else if(ulevel=="Administrator" && level!=wantedlevel){
		tolvl = true;
	}
	
	hasauth := common.CheckUserHasAuthority(a.UUID, ulevel);
	isadmin = hasauth && (a.IsAdmin == common.PreSum("Bdkijw")) && tolvl;
	
	if(!isadmin){
		common.AsyncPrintln("[Security] Verify User as "+ulevel+" has failed!");
		common.AsyncPrintln("[Security]   Username: "+a.UUID);
		common.AsyncPrintln("[Security]   IP: "+a.IP);
	}
	
	return isadmin;
}

// Login the Administrator
func (a Admin) UserLogin(username string, password string) (uverify bool){
	// confirm that this is indeed a true user administrator
	// 	this is IP-locked
	Verifier := &common.VerifyAdmin{Username:a.Username, Password:common.PassSum(common.PreSum(password))};
	common.GetNATSClientJSON().Request("Auth.Admin.Login", Verifier, &a, 100*time.Second);
	
	uverify = a.VerifyUser("Administrator");
	uverify = uverify || a.VerifyUser("VideoModerator");
	uverify = uverify || a.VerifyUser("CommentModerator");
	
	if(!uverify){
		common.AsyncPrintln("[Security] Attempt to login admin failed!");
		common.AsyncPrintln("[Security]   Username: "+a.Username);
		common.AsyncPrintln("[Security]   IP: "+a.IP);
	}
	
	return;
}

