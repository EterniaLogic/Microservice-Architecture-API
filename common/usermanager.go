package common

// User: Brent Clancy (EterniaLogic)
// Date: 1/8/2016

import "time"
import "fmt"
import "container/list"
//import "net/http"

var loginListeners *list.List;
var logoutListeners *list.List;
var registerListeners *list.List;

func InitUserManager(){
	loginListeners = list.New();
	logoutListeners = list.New();
	registerListeners = list.New();

	if(GetNATSClientJSON() == nil){
		panic("InitUserManager: NATS server has not been started!");
	}
	
	GetNATSClientJSON().Subscribe("Auth.Login", func(subj string, reply string, msg *AuthLogin) {
		// somebody has logged in!
		// checks to make sure that UUID is stored
		// 	and that the UUID is stored with it's AuthLevel and TTL
		go UserLoginCheck(msg);
	});
	
	GetNATSClientJSON().Subscribe("Auth.Logout", func(subj string, reply string, msg *interface{}) {
		go UserLogoutCheck((*msg).(string));
	});
	
	GetNATSClientNoJSON().Subscribe("Auth.Register", func(subj string, reply string, msg *interface{}) {
		go UserRegisterCheck((*msg).(string));
	});
}

func CheckDBUserLoginTTL(dofunct interface{}){
	go func(){
		for true {
			time.Sleep(5*time.Second);
			
			rowz,err := db.Query("SELECT HEX(UUID) FROM `"+GetConfig().DBLoginTable+"` WHERE `TTL`<? AND `TTL`>0;", time.Now().Unix());
			if(err != nil){
				// error!
				AsyncPrintln("[Usermanager] TTL Query error");
			}else{
				// loop through rows and add to the list
				for rowz.Next() {
					var UUID string;
					rowz.Scan(&UUID);
					dofunct.(func(string))(UUID);
				}
				rowz.Close();
			}
		}
	}();
}

// listens for a user login
// usage: ListenToLogin(func (ID string){/*stuff*})
func ListenToLogin(f interface{}){
	loginListeners.PushBack(f);
}

// listens for a user logout
// usage: ListenToLogout(func (ID string){/*stuff*})
func ListenToLogout(f interface{}){
	logoutListeners.PushBack(f);
}

// listens for a user register
// usage: ListenToRegister(func (ID string){/*stuff*})
func ListenToRegister(f interface{}){
	registerListeners.PushBack(f);
}

func GetIDFromToken(Token string)(output string){
	errx := GetDB().QueryRow("SELECT HEX(UUID) FROM `"+GetConfig().DBLoginTable+"` WHERE `auth_token`=UNHEX(?);",Token).Scan(&output);
	
	queryq := "SELECT HEX(UUID) FROM `"+GetConfig().DBLoginTable+"` WHERE `auth_token`=UNHEX('"+Token+"');";
	DBAsyncPrintln(queryq);
	
	// check the size
	if(errx != nil){
		//common.AsyncPrintln("GetRowColumnWhere: ERROR=",errx);
		DBAsyncPrintln("GetRowColumnWhere: ERROR="+fmt.Sprintf("%s", errx));
	}
	return output;
}

func GetTokenFromID(ID string)(string){
    token,_ := GetRowColumn(GetConfig().DBLoginTable, ID, "HEX(auth_token)","UUID");
	return token;
}

// Check if the user is logged into this server,
//	if not, add in... else update TTL
func UserLoginCheck(loginmsg *AuthLogin){
	AsyncPrintln("UserManager: User login - "+(*loginmsg).ID);
	// add this new user to the server!
	ID := (*loginmsg).ID;
	if(IsUserLoggedIn(ID)){		
		GetDB().Exec("UPDATE `"+GetConfig().DBLoginTable+"` SET `auth_token`=UNHEX('"+GenRandString()+"') WHERE `UUID`=UNHEX(?)",ID);
	}
	
	// write to map
	//SetMapIndex(ID,loginmsg);
	GetDB().Exec("INSERT INTO `"+GetConfig().DBLoginTable+"` (UUID,username,auth_token,auth_level_gen) VALUES (UNHEX(?),?,UNHEX(?),UNHEX(?))"+
						"ON DUPLICATE KEY UPDATE `auth_token`=UNHEX(?), `auth_level_gen`=UNHEX(?)",
							loginmsg.ID,loginmsg.Username,loginmsg.Token,loginmsg.HashedAuthLevel,
							loginmsg.Token,loginmsg.HashedAuthLevel);
		
	// go through login listener
	for e := loginListeners.Front(); e != nil; e = e.Next() {
		//e.Value((*msg).ID);
		go e.Value.(func(string))(ID);
	}
}

func UserLogoutCheck(ID string){
	AsyncPrintln("UserManager: User logout - "+ID);
	// remove this user from the server
	if(IsUserLoggedIn(ID)){	
		GetDB().Exec("UPDATE `"+GetConfig().DBLoginTable+"` SET `auth_token`=UNHEX('"+GenRandString()+"') WHERE `UUID`=UNHEX(?)",ID);
	}
	
	// go through logout listener
	for e := logoutListeners.Front(); e != nil; e = e.Next() {
		go e.Value.(func(string))(ID);
	}
}

func UserRegisterCheck(ID string){
	AsyncPrintln("UserManager: User register - "+ID);
	// remove this user from the server
	
	// go through logout listener
	for e := registerListeners.Front(); e != nil; e = e.Next() {
		go e.Value.(func(string))(ID);
	}
}

// Checks if the specified user can do something at that level
func CheckUserHasAuthority(ID string, WantedLevel string)(bool){
	b := false;
	if(IsUserLoggedIn(ID)){
		authgen,_ := GetRowColumn(GetConfig().DBLoginTable, ID, "auth_level_gen","UUID")
		b=ConfirmUserLevel(ID, WantedLevel, authgen);
	}
	return b;
}

// returns true if a user is logged in
func IsUserLoggedIn(ID string)(bool){
	id,err := GetRowColumn(GetConfig().DBLoginTable, ID, "UUID","UUID");
	return id == "" || err;
}