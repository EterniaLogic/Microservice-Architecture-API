package common

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015


import "github.com/nats-io/nats"
import "log"
import "time"

// https://github.com/nats-io/nats

// we can use a client-side key for the gnatsd server, but has not been implemented yet.

var conn *nats.EncodedConn; // json
var nconn *nats.Conn; // no json

func StartNATSClient(servers []string, username string, pass string) (*nats.EncodedConn){
	AsyncPrintln("Starting NATS client");
	nc := SecureConnector(servers);
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER);
	
	conn = c;
	nconn = nc;
	return c;
}

func StartNATSClientCFG(){
	StartNATSClient(GetConfig().NATSCluster,
						GetConfig().NATSUser,
						GetConfig().NATSPass);	
}

func SecureConnector(serverURLs []string)(*nats.Conn){
	opts := nats.DefaultOptions;
	opts.Servers = serverURLs;
	opts.Secure = true;
	nc, err := opts.Connect();
	if err != nil {
		log.Fatalf("failed to connect to the NATS messaging server")
	}
	return nc;
}

// send confirmation for user level to the server
func ConfirmUserLevel(id string, wantedLevel string, hashLevel string)(bool){
	// an ID is a UUID without dashes
	mg := &AuthVerifyLevelMsg{ID:id,Wanted:wantedLevel,Hashed:hashLevel}
	mgr := &AuthVerifyLevelRetMsg{ID:id,Authorized:true}
	
	conn.Request("Auth.VerifyLevel", mg, mgr, 10*time.Second);
	return mgr.Authorized;
}

func GetNATSClientJSON()(*nats.EncodedConn){
	return conn;
}

func GetNATSClientNoJSON()(*nats.EncodedConn){
	return conn;
}

func CloseNATSClient(){
	conn.Close();
}