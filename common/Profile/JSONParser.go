package Profile

import "encoding/json"
import "../../common"

type profile struct{
	Email string;
	Firstname string;
	Lastname string;
	Gender string;
	Picture string;
	Bio string;
	City string;
	State string;
	Country string;
	Facebook string;
	Twitter string;
	Skype string;
	Whatsapp string;
	Snapchat string;
	Instagram string;
	Kik string;
	Website string;
	Gear string;
	Birthday string;
}

// Retrieves a user's information based on id
// AuthFunc: Function to call to get information
// idd: interface with id value
// typed: JSON return "Type"
func JSONGetUserProfile(idd interface{})(response []byte){
    id := idd.(string);
    
    // get the username
	var err bool;
    _,err = common.GetRowColumn(common.GetConfig().DBTable,id,"UUID","UUID");
    if(err) { 
        // package error and set JSON
        data := common.MsgData{Type: "Failure", Message: "ID does not exist"};
        tmp, err2 := json.Marshal(data);
        if err2 != nil {
            common.AsyncPrintln("JSON error (JSONGetUserProfile)");
        }
        
        response  = tmp;
    } else { 
        // package data and set JSON
		data := profile{};
		err := common.GetDBGorp().SelectOne(&data, "SELECT email,firstname,lastname,gender,picture,bio,city,state,country,facebook,twitter,skype,whatsapp,snapchat,instagram,kik,website,gear,birthday FROM `"+common.GetConfig().DBTable+"` WHERE `UUID`=UNHEX(?);", id);
		
		if(err == nil){
			tmp, err2 := json.Marshal(data);
			if err2 != nil {
				common.AsyncPrintln("JSON package error (JSONGetUserProfile)");
			}
			
			response  = tmp;
		}else{
			common.AsyncPrintln("SQL error (JSONGetUserProfile)");
		}
		
    }
    return;
}