// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

package main

import "../../common"
import "../../common/Administration"
import "../../common/Auth"
import "../../common/Comment"
import "../../common/Feedback"
import "../../common/Feeds"
import "../../common/Profile"
import "../../common/Search"
import "../../common/Video"
//import "time"
import "fmt"

func main() {
	// console input command scanner
    scanner := common.InitCommandScanner();
	// Add a command to the command scanner for console input
	scanner.AddCommand("help", HelpCommand);
	
	// Configuration file
	common.SetConfig("conf.json");
	
	// onclose function for Ctrl-C or SIGTERM
	common.OnClose();
	
	// based on the wanted config for the server type, select a server to run:
	func(scanner *common.CommandScanner){ 
		switch(common.GetConfig().ServerType){
			case "Admin":
					Administration.Start(scanner);
				break;
			case "Advertising":
					//Advertising.Start(scanner);
				break;
			case "Auth":
					Auth.Start(scanner);
				break;
			case "Comment":
					Comment.Start(scanner);
				break;
			case "Feedback":
					Feedback.Start(scanner);
				break;
			case "Feeds":
					Feeds.Start(scanner);
				break;
			case "Profile":
					Profile.Start(scanner);
				break;
			case "Search":
					Search.Start(scanner);
				break;
			case "Video":
					Video.Start(scanner);
				break;
		}
	}(scanner);
	
	// sleep for a bit before accepting commands
	//time.Sleep(500*time.Millisecond);
	//scanner.ScanLoop();
}

func HelpCommand([]string){
	fmt.Println("Commands:");
	fmt.Println("  exit, quit    - Exit program softly");
	fmt.Println("  clear         - Clears the terminal window");
}
