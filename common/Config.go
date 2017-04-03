package common
// http://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go

import (
    "encoding/json"
    "os"
    "log"
)

type Configuration struct{
	ServerType string;
	_Types string; // not used
	
	DBServer string;
	DB string;
	DBTable string;
	DBLoginTable string;
	DBUser string;
	DBPass string;
	
	NATSCluster []string;
	NATSUser string;
	NATSPass string;
	
	RESTfulPort string;
	
	_Comment string; // only a comment
	LoginTTL int;
	MailURL string;
	MailKey string;
	DoEmail bool;
}

var config Configuration;

func SetConfig(fileloc string){
	file, _ := os.Open(fileloc)
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
	  log.Fatalf("error:", err)
	}
}

func GetConfig() Configuration{
	return config;
}