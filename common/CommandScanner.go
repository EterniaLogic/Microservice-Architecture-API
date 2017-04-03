package common

import "fmt"
import "strings"
import tm "github.com/buger/goterm"
import "os"
import "os/exec"

type CommandScanner struct{
	comlist map[string]func([]string);
};
type CommandExec struct{
	command string; // command
	funct interface{}; // function to execute
};


func InitCommandScanner() *CommandScanner{
	scanner := &CommandScanner{};
	scanner.comlist = make(map[string]func([]string));
	scanner.AddCommand("exit", ExitCommand);
	scanner.AddCommand("quit", ExitCommand);
	scanner.AddCommand("clear", ClearCommand);
	
	cmd := exec.Command("clear");
	cmd.Stdout = os.Stdout;
	cmd.Run();
	
	tm.Clear(); // clear current screen
	tm.MoveCursor(1,1);
	tm.Flush();
	return scanner;
}

func (c CommandScanner) ScanLoop(){
	for(true){
		fmt.Print("> ");
		var response string;
		_, err := fmt.Scanln(&response);
		if(err != nil){
			//fmt.Println(err);
		}
		
		if(strings.Index(response," ") == -1){
			if(c.comlist[response] != nil){
				c.comlist[response]([]string{response});
			}else{
				fmt.Println("Unknown command");
			}
		}else{
			Strings := strings.Split(response, " ");
			if(c.comlist[Strings[0]] != nil){
				c.comlist[Strings[0]](Strings);
			}else{
				fmt.Println("Unknown command");
			}
		}
		
	}
}

func (c CommandScanner) AddCommand(com string, funct func([]string)){
	c.comlist[com] = funct;
}

func ExitCommand(args []string){
	OnExit();
}

func ClearCommand(args []string){
	tm.Clear();
	tm.MoveCursor(1,1);
	tm.Flush();
}