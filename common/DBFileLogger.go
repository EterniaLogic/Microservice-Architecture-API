package common

import (
	"os"
	"log"
	"time"
)

var channeldb chan string;
// The DBFileLogger is a asynchronous logger that allows for better performance
//	by preventing blocking code.

// "DB" is for debugging database code.

var filedb *os.File;
func LogToDBFile(fileloc string, dofile bool){
	channeldb = make(chan string);
	go DBAsyncLogger(fileloc, dofile);
}

func DBAsyncPrintln(i string){
	channeldb <- i;
}

// thread that logs to file over time
func DBAsyncLogger(fileloc string, dofile bool){
	log.Println("Logging DB to",fileloc);
		
	f, _ := os.OpenFile(fileloc, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0600);
	dblogger := log.New(f, "logger: ", log.Lshortfile);
	filedb = f;

	for true {
		for data := range channeldb {
			dblogger.Println(data);
			time.Sleep(1*time.Millisecond);
		}
		
		time.Sleep(1*time.Second);
	}
}

// close logging to file at the end of the program
func CloseDBLogToFile(){
	if(file == nil) { return ; }
	filedb.Sync();
	filedb.Close();
}