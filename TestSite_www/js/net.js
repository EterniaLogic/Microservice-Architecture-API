// Networking Javascript file
// © Brent Clancy 2017
// Used to connect with the API
// Author: Brent Clancy (3/1/2016)

// Unoptimized v0.1

// List of API server locations
var APILoc = "http://eternialogic.com:630";

var APIP = "/api/v1";
var AuthLoc = APIP+"/auth";
var CommentsLoc = APIP+"/comments";
var FeedbackLoc = APIP+"/feedback";
var FeedsLoc = APIP+"/feeds";
var ProfilesLoc = APIP+"/profiles";
var SearchLoc = APIP+"/search";
var VideosLoc = APIP+"/videos";

Token = "";
UUID = "";

function Getxhttp(){
	var xhttp;
	if (window.XMLHttpRequest) {
		xhttp = new XMLHttpRequest();
		} else {
		// code for IE6, IE5
		xhttp = new ActiveXObject("Microsoft.XMLHTTP");
	}
	return xhttp;	
}

function getxHttpData(method,url,data){
	var xhttp = Getxhttp();
	
	xhttp.open(method, url, true);
	xhttp.send(data);
	
	return xhttp;
}

function Login(username,password,contentdiv,responsediv){
	var data = '{"Username":"'+username+'","UserPass":"'+password+'"}';
	var xhttp = getxHttpData("POST",APILoc+AuthLoc+"/login",data);
	
	xhttp.onreadystatechange = function() {
	  if (xhttp.readyState == 4 && xhttp.status == 200) {
		onLogin(contentdiv,responsediv,xhttp.responseText);
	  }
	};
}

function onLogin(contentdiv,responsediv,response){
	console.log(responsediv);
	document.getElementById(responsediv).innerHTML = response;
	
	// parse and use json
	var obj = JSON.parse(response);
	if(obj !== null){
		if(obj.Type !== null){
			if(obj.Type == "Authentication") {
				Token = obj.Token;
				UUID = obj.ID;
				destroyWindow(contentdiv);
				
				SetCookie(_cookieNames[0],Token);
				SetCookie(_cookieNames[1],UUID);
			}
		}
	}
}


