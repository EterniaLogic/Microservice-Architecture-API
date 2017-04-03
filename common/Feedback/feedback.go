package Feedback

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

import "../../common"

func Start(scanner *common.CommandScanner){
	common.LogToFile("FeedbackLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Feedback Server...");
	
	// open a DB
    common.OpenDBCFG();
	
	// connect to NATS messaging server
	common.StartNATSClientCFG(); // between-microservices data
	common.InitUserManager(); // handles users (usermanager.go)
	
    // wait for aggregate server to return info
    StartRESTfulServer();
}