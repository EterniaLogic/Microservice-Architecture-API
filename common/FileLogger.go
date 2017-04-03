package common

import (
	"os"
	"log"
	"time"
)

var channel chan string;

// The FileLogger is a asynchronous logger that allows for better performance
//	by preventing blocking code.

var file *os.File;
func LogToFile(fileloc string, dofile bool){
	channel = make(chan string);
	go AsyncLogger(fileloc, dofile);
}

func AsyncPrintln(i string){
	channel <- i;
}

// thread that logs to file over time
func AsyncLogger(fileloc string, dofile bool){
	if(dofile){
		log.Println("Logging to",fileloc);
		f, _ := os.OpenFile(fileloc, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0600);
		log.SetOutput(f);
		file = f;
	}

	for true {
		for data := range channel {
			log.Println(data);
			time.Sleep(1*time.Millisecond);
		}
		
		time.Sleep(1*time.Second);
	}
}

// close logging to file at the end of the program
func CloseLogToFile(){
	if(file == nil) { return ; }
	file.Sync();
	file.Close();
}