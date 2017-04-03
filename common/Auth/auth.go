package Auth

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015


//import "golang.org/x/oauth2/facebook"
//import "golang.org/x/oauth2/google"
import "database/sql"
import "../../common"
import "time"
import "strings"
import "math"

func Start(scanner *common.CommandScanner){
	common.LogToFile("AuthLog.txt",true);
	common.AsyncPrintln("[INIT] Starting Authorization Server...");
	
	// open a DB
	common.OpenDBCFG();
	common.InitMailer();
	
	common.StartNATSClientCFG(); // between-microservices data
	InitSubscribeNAT();
	
	// User Session manager with TTL checking
	common.CheckDBUserLoginTTL(AuthLogout);
	
	// wait for aggregate server to return info
	StartRESTfulServer();
}


// Authenticate a Direct user
// username, password direct
// returns:
//      UserID - database id (1-N)
//      idtoken - Passed to client for seamless pages
//      msg - eror message output
//      err - true if an error has occured
func AuthUser(username string, passwd string, prehash bool)(id string, idtoken string, msg string, err bool){
    common.AsyncPrintln("Auth User: "+username);
	
	if(prehash){
		passwd = common.PreSum(passwd);
	}
    
    // compare username & password
    // check the user if it exists
    if(AuthUserExists(username)){
        // The user exists!
		
		var UUID, password, auth_level, auth_level_gen string;
		common.GetDB().QueryRow("SELECT HEX(UUID),HEX(password),auth_level,HEX(auth_level_gen) FROM `"+common.GetConfig().DBTable+"` WHERE `username`=?;",username).Scan(&UUID,&password, &auth_level, &auth_level_gen);
        
        // determine if the password matches with the user:
        if(strings.ToUpper(common.PassSum(passwd)) == strings.ToUpper(password)){
            // Check/fix the authorization level
            auth_level, auth_level_gen = CheckAuthLevel(UUID, auth_level, auth_level_gen);
            
            // Additionally, determine if the user is banned or has not replied to their register email
            if(common.ConvertStringToLevel(auth_level) < 5){
				if(common.ConvertStringToLevel(auth_level) == 0){
					msg = "You have been banned";
					common.AsyncPrintln("AuthUser: Banned user login ("+username+")");
				}else{
					msg = "Have you checked your email yet?";
					common.AsyncPrintln("AuthUser: Unverified email user login ("+username+")");
				}
				err=true;
            }else{
                // login accepted, generate a token and return id
				msg = "Logged in";
                id,idtoken = AuthUserAccepted(UUID, username, auth_level, auth_level_gen);
            }
        }else{
            msg = "Username/Password incorrect";
			common.AsyncPrintln("AuthUser: Password incorrect ("+username+")");
            err=true;
        }
    }else{
        msg = "Username/Password incorrect";
		common.AsyncPrintln("AuthUser: Username incorrect ("+username+")");
        err = true;
    }
    
    return;
}

// id and secret are privided by a website.
// the website is: google.com, facebook.com
// returns:
//      UserID - database id (1-N)
//      idtoken - Passed to client for seamless pages
//      msg - eror message output
//      err - true if an error has occured
func OAuthUser(token string, website string)(idret string, idtoken string, msg string, err bool){
    //var AuthURL, TokenURL string;
    //common.AsyncPrintln("OAuth User");
    
	
	// Check if the user has an OAuth token in their account
    var UUID, Username, password, auth_level, auth_level_gen string;
	common.GetDB().QueryRow("SELECT HEX(UUID),username,HEX(password),auth_level,HEX(auth_level_gen) FROM `"+common.GetConfig().DBTable+"` WHERE `oauth2_token`=?;",token).Scan(&UUID,&Username,&password, &auth_level, &auth_level_gen);
	if(UUID != ""){
		// login accepted, generate a token and return id
		idret,idtoken = AuthUserAccepted(UUID, Username, auth_level, auth_level_gen);
		auth_level, auth_level_gen = CheckAuthLevel(UUID, auth_level, auth_level_gen);
            
		// Additionally, determine if the user is banned or has not replied to their register email
		if(common.ConvertStringToLevel(auth_level) < 5){
			if(common.ConvertStringToLevel(auth_level) == 0){
				msg = "You have been banned";
				common.AsyncPrintln("AuthUser: Banned user login ("+Username+")");
			}else{
				msg = "Have you checked your email yet?";
				common.AsyncPrintln("AuthUser: Unverified email user login ("+Username+")");
			}
			err=true;
		}else{
			// login accepted, generate a token and return id
			idret,idtoken = AuthUserAccepted(UUID, Username, auth_level, auth_level_gen);
		}
		
	}else{
        msg = "Bad OAuth token";
		common.AsyncPrintln("AuthUser: Oauth2 incorrect ("+token+")");
        err = true;
    }
    
    return;
}


// id and secret are privided by a website.
// the website is: google.com, facebook.com
// requires:
//		UUID - user id string
// returns:
//      idtoken - generated token
//		id - UUID
func AuthUserAccepted(UUID string, username string, auth_level string, auth_level_gen string) (id string, idtoken string){
	idtoken = common.GenRandString();
	id=UUID;
	go func(){
		// update LastLogin on table
		nowTime := int(time.Now().Unix());
		nowTime = nowTime+common.GetConfig().LoginTTL;
		common.GetDB().Exec("UPDATE "+common.GetConfig().DBTable+" SET `last_login`=NOW(), `auth_token`=UNHEX(?), `TTL`=? WHERE UUID=UNHEX(?)",idtoken,nowTime,UUID);
		
		// send NATS message for login
		loginmsg := common.AuthLogin{ID:UUID,Username:username,Token:idtoken,HashedAuthLevel:auth_level_gen};
		common.GetNATSClientJSON().Publish("Auth.Login",loginmsg);
	}();	
	
	return;
}


func AuthUserExists(uname string) bool{
	var username string;
    errx:=common.GetDB().QueryRow("SELECT username FROM "+common.GetConfig().DBTable+" WHERE username=?;",uname).Scan(&username);
    
    // check the size
    if(errx != sql.ErrNoRows || username == uname){
		return true;
	}else{
		return false;
	}
}


// Create a new user
// returns:
//      uname - username to create
//      password - plaintext or hashed password
//      msg - eror message output
//      err - true if an error has occured
func AuthRegister(uname string, password string, email string)(msg string, err bool){
    // Register a user
    
    passhashed := common.PassSum(password);
    // check the size
    if(AuthUserExists(uname)){
        msg = "User already exists";
		common.AsyncPrintln("AuthRegister: Username already exists ("+uname+")");
        err = true;
    }else{
        common.AsyncPrintln("User register: "+uname+" "+email);
        common.GetDB().Exec("INSERT INTO "+common.GetConfig().DBTable+" (UUID,username,password,auth_level,join_date) VALUES (UNHEX(REPLACE(UUID(),'-','')),?,UNHEX(?),?,NOW());",
                        uname,passhashed,"Unverified");
		
		var UUID string;
		common.GetDB().QueryRow("SELECT HEX(UUID) FROM "+common.GetConfig().DBTable+" WHERE username=?;",uname).Scan(&UUID);
		
		authlgen := common.GenAuthLevel("Unverified",UUID);
		common.GetDB().Exec("UPDATE "+common.GetConfig().DBTable+" SET `auth_level_gen`=UNHEX(?) WHERE `UUID`=UNHEX(?);",authlgen,UUID);
		
		common.GetNATSClientNoJSON().Publish("Auth.Register",UUID);
		
		msg = "User registered, please check your email to verify your account.";
		
		// send email for verification
		AuthSendVerificationEmail(UUID,email);
    }
    
    return;
}


// Create a new user
// Inputs:
//      UUID - User ID
//      email - unverified user to send email to
func AuthSendVerificationEmail(UUID string, email string){
	mail := common.GetMailer();
	
	// generate a token for registration
	regtoken := common.GenRandString();
	common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"` SET `auth_token`=UNHEX(?) WHERE `UUID`=UNHEX(?)",regtoken,UUID);
	
	m := mail.NewMessage(  
		"Noreply <noreply@"+common.GetConfig().MailURL+">", // From
		"User Registration",                    // Subject
		"User account verification: http://api.site.com/new-user-verification/"+regtoken, // Plain-text body
		"<"+email+">");        	// Recipients (vararg list)
	
	m.SetHtml("User account verification: <br /><a href='http://api.site.com/new-user-verification/"+regtoken+"'>Click here to finish applying!</a>");

	if(common.GetConfig().DoEmail){
		mail.Send(m);
	}
}


// Verify a user's level
// returns:
//      id - uid from the original login
//      wanted - level that the user wants ie: "Administrator", "User"
//      hashed - hashed user level that is given with GetUserLevel
//      msg - output error if UID doesnt exist
//      err - true if an error has occured
func AuthVerifyUserLevel(id string, wanted string, hashed string)(msg string, err bool){
    var authlv, authgen string;
	var err1 bool;
    authlv,err1 = common.GetRowColumn(common.GetConfig().DBTable,id,"auth_level","UUID");
    authgen,_ = common.GetRowColumn(common.GetConfig().DBTable,id,"HEX(auth_level_gen)","UUID");
    
    //common.AsyncPrintln(username,authlv,authgen);
    
    if(err1 == true){
        err = true;
        msg = "User ID not found";
		common.AsyncPrintln("AuthVerifyUserLevel: User ID for "+wanted+" does not exist ("+id+")");
    }else{
        // get verified Auth Level of the user on the tables
        // if it is not correct, auto-patch it to "User" on the database
		CheckAuthLevel(id, authlv, authgen);
		
        level := common.ConvertStringToLevel(authlv);
        wantedlevel := common.ConvertStringToLevel(wanted);
        
		if(wanted=="CommentModerator" && math.Mod(float64(level),2)!=0){ // CommentMod
			msg = "false";
            err = true;
		}else if(wanted=="VideoModerator" && math.Mod(float64(level),3)!=0){
			msg = "false";
            err = true;
		}else if(wanted=="Administrator" && level!=wantedlevel){
			msg = "false";
            err = true;
		}
    }
    return;
}


// AuthLogout will log a user out
func AuthLogout(id string){
    common.AsyncPrintln("User logout: "+id);
    go func(){
		// clear the token for this user
		common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"` SET `auth_token`='',TTL=0 WHERE `UUID`=UNHEX(?);",id);
		
		// send NATS response to all other microservices
		common.GetNATSClientJSON().Publish("Auth.Logout",id);
	}();
}

// get verified Auth Level of the user on the tables
// if it is not correct, auto-patch it to "User" on the database
// returns:
//      id - generated
//      authlv - 
//      authgen - 
// outputs:
//		new authlv (if changed)
//		new authgen (if changed)
func CheckAuthLevel(id string, authlv string, authgen string)(string, string){
    auth := common.GenAuthLevel(authlv, id);
    if(strings.ToUpper(auth) != strings.ToUpper(authgen)){
        // auto patch on DB
		common.AsyncPrintln("CheckAuthLevel: SECURITY WARNING! UID: "+id+" was trying for a different auth level: '"+authlv+"' on the database. Modified SQL field?");
		
        authgen = common.GenAuthLevel("Unverified", id);
        common.GetDB().Exec("UPDATE `"+common.GetConfig().DBTable+"` SET `auth_level`=?, `auth_level_gen`=UNHEX(?) WHERE `UUID`=UNHEX(?)","Unverified",authgen,id);
    }
	
	return authlv, authgen;
}