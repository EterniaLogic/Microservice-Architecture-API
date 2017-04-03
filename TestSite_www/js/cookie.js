// Cookie jar Javascript file
// © Brent Clancy 2017
// Manages active cookies
// Author: Brent Clancy (3/1/2016)

_cookieNames = ["tk","u"];

function SetCookie(cookiename, value){
	document.cookie = cookiename+"="+value;
}

function GetCookie(cookiename){
	
}