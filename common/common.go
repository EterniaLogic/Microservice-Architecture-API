package common

import "os"
import "os/signal"
import "syscall"

func OnClose(){
	// end of loop, close out
	// these are executed upon closing
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func(){
		<-c;
		OnExit();
	}()
}

func OnExit(){
	CloseDB();
	CloseNATSClient(); // close messaging client connection
	AsyncPrintln("Closing Server...");
	CloseLogToFile();
	os.Exit(0);
}