// Video Javascript file
// © Brent Clancy 2017
// Used to Get/Post videos
// Author: Brent Clancy (3/1/2016)

var searchWindow=null;

// Draw a login window
function DrawAddVideoWindow(lat,lng){
	var width = 300;
	var height = 80;
	
	var content = 'Find Video by curator: <span class="Link" onclick="helpWindow(\'FindVideo\');">(?)</span><br />';
	content += '<img src="https://www.youtube.com/yt/brand/media/image/yt-brand-flattype-2.png" class="IconYoutube" onclick="YoutubeWindowVid();" />';
	content += '<img src="img/logos/vimeo_logo_blue.png" class="IconVimeo" onclick="VimeoWindowVid();" />';
	content += '<div id=\"VidData\"></div>';
	content += '<input type="hidden" id="VidLoc" value="'+lat+' '+lng+'" />&nbsp;&nbsp;<input type="hidden" id="VidId" value="" /><input type="hidden" id="VidWebsite" value="" /><br />';
	//content += '<br />Video Description:<Br />&nbsp;&nbsp;<textarea onclick="this.select();" style="width:280px;height:60px" id="VidDesc"></textarea><div align=right width=300><input type="button" id="VidButton" disabled=true value="Add Video!" />&nbsp;</div>';
	var windowt = drawWindow("Add Video",content,document.body.clientWidth/2-width/2,35,width,height);
	document.getElementsByTagName('body')[0].appendChild(windowt[0]);
	return windowt;
}

function SelectVideoAdd(Id,title,thumb,desc,website){
	var content = "<hr /><table><tr><td><a href='"+website+Id+"' target='_BLANK'><img src='"+thumb+"' /></a></td><td valign=top><b>"+title+"</b></td></tr></table><hr />";
	content += 'Video Description:<Br />&nbsp;&nbsp;<textarea onclick="this.select();" style="width:280px;height:60px" id="VidDesc"></textarea><div align=right width=300><input type="button" id="VidButton" disabled=true value="Add Video!" />&nbsp;</div>';
	
	
	document.getElementById("VidData").innerHTML = content;
	document.getElementById("VidDesc").innerHTML = desc.replace("&quot;","\"");
	document.getElementById("VidId").value = Id;
	document.getElementById("VidWebsite").value = Id;
	document.getElementById("VidButton").disabled=false;
	
	// resize window
	document.getElementById("VidWebsite").parentNode.parentNode.style.height = (document.getElementById("VidWebsite").parentNode.offsetHeight+25)+"px";
	
	searchWindow[0].parentNode.removeChild(searchWindow[0]);
	searchWindow=null;
}