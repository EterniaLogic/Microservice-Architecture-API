// UI Javascript file
// © Brent Clancy 2017
// Used to manipulate the sidebar and added animations
// Author: Brent Clancy (3/1/2016)

var toggledMenu = false;
var menuwidth = 1; // current width of menu
var speed = 13; // speed for menu open/close
var time = 0.5;
var change = time/(speed/1000);
var winInterval; // animation interval
var menuid = "sidebar";
var animating = false;
var selectedSidebar = 0;

function isMenuOpen(){
	return toggledMenu;
}

function toggleMenu(){
	// menu location
	var mobile = detectMobile();
	toggledMenu = !toggledMenu;
	
	if(winInterval != null) clearInterval(winInterval);
	winInterval = setInterval("animateMenu("+mobile+");",speed);
}

function animateMenu(mobile){
	// Figure out what width we need to go to:
	var targetx = toggledMenu ? (mobile ? document.body.clientWidth : document.body.clientWidth*0.4) : 1;
	var delta = toggledMenu ? change : -change;
	
	if((menuwidth <= targetx && !toggledMenu) || (targetx <= menuwidth && toggledMenu)){
		menuwidth = targetx;
		clearInterval(winInterval);
	}else{
		menuwidth = (menuwidth+delta);
	}
	
	//console.log(toggledMenu+" "+targetx+" "+menuwidth+" "+delta);
	document.getElementById(menuid).style.width=menuwidth+"px";
}

function sidebarItem(id,this1){
	var SidebarData = document.getElementById("sidebar-content");
	
	// replace coloring
	getSidebarItem(selectedSidebar).className = "sidebar-item";
	this1.className = "sidebar-selected";
	selectedSidebar=id;
	
	switch(id){
		case 0:
				document.getElementById("sidebar-content").innerHTML = sbGetFeeds();
			break;
		case 1:
				document.getElementById("sidebar-content").innerHTML = sbGetRecent();
			break;
		case 2:
				document.getElementById("sidebar-content").innerHTML = sbGetAccount();
			break;
		case 3:
				document.getElementById("sidebar-content").innerHTML = sbGetFeedback();
			break;
		case 4:
				document.getElementById("sidebar-content").innerHTML = sbGetAbout();
			break;
		default:
		
			break;
	}
}

function getSidebarItem(id){
	return document.getElementById("sb"+id);
}

// https://developers.google.com/maps/documentation/javascript/basics
function detectMobile() {
  var useragent = navigator.userAgent;

  if (useragent.indexOf('iPhone') != -1 || useragent.indexOf('Android') != -1 ) {
    return true;
  } else {
    return false;
  }
}