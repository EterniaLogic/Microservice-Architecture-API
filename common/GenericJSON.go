package common

import "encoding/json"
import "time"
import "fmt"

// User: Brent Clancy (EterniaLogic)
// Date: 12/23/2015

// convert a json string to a string map interface
func JSONtoMap(JSON string)(out map[string]interface{},err bool){
    var data interface{};
    errx := json.Unmarshal([]byte(JSON), &data);
    
    // if json was formatted correctly
    if(errx == nil){
        out = data.(map[string]interface{});
	}else{
		fmt.Println("JSONtoMap",errx);
		err = true;
	}
	return;
}

func GetJSONMapString(inmap map[string]interface{}, column string)(string, bool){
	if val, ok := inmap[column]; ok {
		return val.(string),false;
	}else{
		AsyncPrintln("GetJSONMapString ERROR: Attempt to get row more than map allows");
		return "",true;
	}
}


// Retrieves a user's information based on id
// AuthFunc: Function to call to get information
// idd: interface with id value
// typed: JSON return "Type"
func JSONGetIDColumn(table string, idd interface{}, sqlcolumn string, typed string, idcol string)(response []byte){
    id := idd.(string);
    
    // get the data
    data,err := GetRowColumn(table, id, sqlcolumn, idcol);
    if(err) { 
        // package error and set JSON
        data := MsgData{Type:"Failure", Message: "ID does not exist"};
        tmp, err2 := json.Marshal(data);
        if err2 != nil {
            AsyncPrintln("JSON error (JSONGetIDColumn)");
        }
        
        response  = tmp;
    } else { 
        // Postformatting for LastLogin to human readable date and time
        if(sqlcolumn == "lastlogin"){
            date,erry := time.Parse("2006-01-02 15:04:05", data);
            if erry != nil {
                AsyncPrintln("JSON Date Conversion error");
            }
            data = date.Format("Jan 2, 2006 15:4:5 PM");
        }
        
        // package data and set JSON
        data := ValueX{Type: typed, Value:data};
        tmp, err2 := json.Marshal(data);
        if err2 != nil {
            AsyncPrintln("JSON marshal error (JSONGetIDColumn)");
        }
        
        response  = tmp;
    }
    
    
    return;
}

// JSON packaging a message that does "Failure", success
// Inputs:
//		successmsg - user's id to be sent
//		msg - taken in from the request, used for an error
//		err - error from authlogin
// Outputs:
//		response - JSON response to the RESTful client
func JSONPackageMessage(successmsg string, msg string, err bool)(response []byte){
    var typed string; 
    if(err) { 
        typed = "Failure"; 
    } else { 
        typed = successmsg; 
    }
    
    // package data and set JSON
    data := MsgData{Type:typed, Message: msg};
    tmp, err2 := json.Marshal(data);
    if err2 != nil {
        AsyncPrintln("JSON error (JSONPackageMessage)");
    }
    
    response  = tmp;
    return;
}

// Tests if an input, such as "Username" is given through 
// the client's input JSON request.
// Inputs:
//		data - data to test if not null
//		msg - output message in JSON response
// Outputs:
//		response - JSON response to the RESTful client
func TestJSONInput(data interface{},msg string)(response interface{}){
    response = nil;
    
    // test the input
    if(data == nil){
		AsyncPrintln("err");
        response = JSONPackageMessage("", msg, true);
    }
    return;
}