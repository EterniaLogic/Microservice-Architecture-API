// Google Javascript file
// © Brent Clancy 2017
// Used to manipulate custom windows
// Author: Brent Clancy (3/1/2016)

GoogleAPI = "";
gdata = null;

function YoutubeSearch(q, callback) {
	var xhttp = getxHttpData("GET","https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=50&key="+GoogleAPI+"&q="+q,"");
	
	xhttp.onreadystatechange = function() {
	  if (xhttp.readyState == 4 && xhttp.status == 200) {
		var data = JSON.parse(xhttp.responseText);
		gdata = data;
		callback(data);
	  }
	};
}

// Show a window to find a Youtube video
function YoutubeWindowVid(){
	//if(detectMobile()){
	var width = document.body.clientWidth-20;
	var height = document.body.clientHeight-31;
	var content = "Youtube Search: <input type='text' id='SearchBox' onclick=\"this.select();\" /><input type=button value='Search' onclick='YoutubeSearch(document.getElementById(\"SearchBox\").value,YoutubeWindowVidCallback);' /><br />";
	content += "<div id='YoutubeSearchBox'></div>";
	
	var windowt = drawWindow("Find Youtube Video",content,5,26,width,height);
	document.getElementsByTagName('body')[0].appendChild(windowt[0]);
	searchWindow = windowt;
	//YoutubeSearch(,);
}

// Callback when searching for videos on Youtube
function YoutubeWindowVidCallback(data){
	var items = data.items;
	var SearchBox = document.getElementById("YoutubeSearchBox");
	var htmlData = "<table>";
		
	// loop through items
	for(var i in items){
		if(items[i].id.kind == "youtube#video"){
			var item = items[i]
			var vidid = item.id.videoId;
			var vidtitle = item.snippet.title;
			var vidthumb = item.snippet.thumbnails.default.url;
			var viddesc = item.snippet.description;
			htmlData += "<tr onclick='SelectVideoAdd(\""+vidid+"\",\""+vidtitle+"\",\""+vidthumb+"\",\""+viddesc+"\",\"https://youtube.com/watch?v=\")'><td style='width:120px'><img src='"+vidthumb+"' /></td>";
			htmlData += "<td><b><h4>"+vidtitle+"</h4></b>by <i><a href='https://www.youtube.com/channel/"+item.snippet.channelId+"' target='_BLANK'>"+item.snippet.channelTitle+"</a></i><br />"+viddesc+"</td></tr>";
		}
	}
	
	htmlData += "</table>";
	
	SearchBox.innerHTML = htmlData;
	SearchBox.style.height = (SearchBox.parentNode.parentNode.style.height.replace("px","")-45)+"px";
}
