// Main Javascript file
// © Brent Clancy 2017
// Author: Brent Clancy (3/1/2016)

function onLoad(){
	console.log("loaded");
	
	// DrawLoginWindow();
	
	document.getElementById("topmenu").onclick=toggleMenu;
	document.getElementById("toplogo").onclick=clickLogo;
	sidebarItem(0,document.getElementById("sb0"));
}



function clickLogo(){
	location.href="http://app.com";
}

