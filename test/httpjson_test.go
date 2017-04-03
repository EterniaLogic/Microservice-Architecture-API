package test

import "net/http"
import "strings"
import "log"
import "testing"
import "io/ioutil"
import "math/rand"
import "strconv"
import "time"

// This will test ALL json/http requests all at once
// All servers MUST be on all at the same time on the same system for this to work.


//req.Header.Add("If-None-Match", `W/"wyzzy"`)
var client = &http.Client{
	
}
var username string;
var Token string; // "gstx"

// begin initializations for the client
func TestJSONInit(t *testing.T){
	rand.Seed(time.Now().UnixNano());
	Token = "";
	username = "testusr_"+strconv.Itoa(rand.Intn(999999999));
}

func TestJSONAuth(t *testing.T){

	// Create user:
	//locTest("POST", "application/json", "http://localhost:6101/api/v1/auth/user", "{\"Username\":\""+username+"\",\"UserPass\":\"test\",\"Email\":\"eternialogic@gmail.com\"}", "{\"Type\":\"Success\"", t);
	
	// login user: (has checked email?)
	//locTest("POST", "application/json", "http://localhost:6101/api/v1/auth/login", "{\"Username\":\""+username+"\",\"UserPass\":\"test\"}", "{\"Type\":\"Failure\",\"Message\":\"Have you checked your email yet?\"}", t);
	
	
	// Verify the user
	// http://api.site.com/api/v1/auth/verify/2C7FC0DB311E26893F37849ACEC264FC1E777A7BB875158A9D4001DA
	//locTest("GET", "application/json", "http://localhost:6101/api/v1/auth/verify/"+response, "", "Success", t);

	
	// login test the user
	tstr := locTest("POST", "application/json", "http://localhost:6101/api/v1/auth/login", "{\"Username\":\"testusr_997112853\",\"UserPass\":\"test\"}", "{\"Type\":\"Authentication\"", t);
	log.Println(tstr);
	
	// Get data from loginstr
	
	// logout user:
	//locTest("DELETE", "application/json", "http://localhost:6101/api/v1/auth/login", "{}", "{\"Type\":\"Authentication\"", t);
}


func TestJSONProfile(t *testing.T){
	
}


func locTest(Type string, contenttype string, loc string, data string, expected string, t *testing.T)(string){
	req, err := http.NewRequest(Type, loc, strings.NewReader(data));
	req.Header.Add("Content-Type",contenttype);
	req.Header.Add("gstx",Token);
	resp, err2 := client.Do(req);
	
	if(err != nil || err2 != nil){
		log.Println(err);
		log.Println(err2);
		t.Fail();
	}
	
	datax,_ := ioutil.ReadAll(resp.Body);
	
	if(!strings.Contains(string(datax),expected)){
		log.Println("Expected data was not matched! ",loc,"  ",data," ",string(datax));
		t.Fail();
	}
	return string(datax);
}