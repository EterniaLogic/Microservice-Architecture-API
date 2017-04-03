// Main Javascript file
// © Brent Clancy 2017
// Sidebar content drawing and networking intermediates
// Author: Brent Clancy (3/1/2016)

function sbGetFeeds(){
	return "feeds";
	// Request feeds from Feeds API
}

function sbGetRecent(){
	return "recent";
	// Request recent from Videos API
}

function sbGetAccount(){
	return "account";
	// Request account data from account API
}

// Generate a feedback form:
function sbGetFeedback(){
	return "feedback";
	// return a standard Feedback form
}

// Show copyright, company info and logo
function sbGetAbout(){
	var about = "<center><img src='img/logos/logolarge.jpg' /></center><br />";
	about += '<a href="http://www.instagram.com/app" target="_BLANK"><span class="fa fa-instagram">&nbsp;&nbsp;</span></a>&nbsp;';
	about += '<a href="http://www.twitter.com/app" target="_BLANK"><span class="fa fa-twitter">&nbsp;&nbsp;</span></a>&nbsp;';
	about += '<a href="http://www.facebook.com/app" target="_BLANK"><span class="fa fa-facebook">&nbsp;&nbsp;</span></a>';
	about += '<br /> App<br />© Brent Clancy 2016 (Pending transfer)';
	
	// Pending transfer to, but not applicable until delivered:
	//about += '<br /> App<br />© Brent Clancy 2017';
	return about;
}
