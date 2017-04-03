package Auth
import "encoding/json"
import "../../common"

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015





// convert a json string to a string map interface
func JSONtoMap(JSON string)(out map[string]interface{},err bool){
    var data interface{};
    errx := json.Unmarshal([]byte(JSON), &data);
    
    // if json was formatted correctly
    if(errx == nil){
        out = data.(map[string]interface{});
	}else{
		err = true;
	}
	return;
}

// logs out a user by taking in an ID response is in JSON
func JSONDeAuth(i interface{})(response []byte){    
	id := i.(string);
	
    // do the logout
    AuthLogout(id);
    response = common.JSONPackageMessage("Success","User logged out",false);
    
    return;
}

// Retrieves a user's information based on id
// AuthFunc: Function to call to get information
// idd: interface with id value
// typed: JSON return "Type"
func JSONGetUserProfile(idd interface{})(response []byte){
    id := idd.(string);
    
    // get the username
    var col1,col2,col3 string;
	var err bool;
    col1,err = common.GetRowColumn(common.GetConfig().DBTable,id,"realname","UUID");
    col2,_ = common.GetRowColumn(common.GetConfig().DBTable,id,"lastlogin","UUID");
    col3,_ = common.GetRowColumn(common.GetConfig().DBTable,id,"auth_level","UUID");
    if(err) { 
        // package error and set JSON
        data := common.MsgData{Type:"Failure", Message: "ID does not exist"};
        tmp, err2 := json.Marshal(data);
        if err2 != nil {
            common.AsyncPrintln("JSON error (JSONGetUserProfile)");
        }
        
        response  = tmp;
    } else { 
        // package data and set JSON
        data := common.UserProfile{Type: "UserProfile", RealName:col1, LastLogin:col2, UserLevel:col3};
        tmp, err2 := json.Marshal(data);
        if err2 != nil {
            common.AsyncPrintln("JSON packaging error (JSONGetUserProfile)");
        }
        
        response  = tmp;
    }
    return;
}

// Used to output login verification and fail modes to the RESTful client
// Inputs:
//		id - user's id to be sent
//		token - user's token to be sent
//		msg - taken in from the AuthLogin request
//		err - error from authlogin
// Outputs:
//		response - JSON response to the RESTful client
func JSONPackageLogin(id string, token string, msg string, err bool)(response []byte){
    var data interface{};
    if(err) { // An error has occured, send a message
        // package data and set JSON
        data = common.MsgData{Type:"Failure", Message: msg};
    } else { 
        // Successful login, send data
        data = common.Authentication{Type:"Authentication", ID: id, Token: token};
    }
    
    tmp, err2 := json.Marshal(data);
    if(err2 != nil){
        common.AsyncPrintln("JSON packaging error (JSONPackageLogin)");
    }
    response = tmp;
    
    return;
}