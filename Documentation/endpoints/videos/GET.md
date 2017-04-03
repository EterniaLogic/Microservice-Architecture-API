# <code>GET</code> /api/v1/videos/:vid

## Token Required
	none or Header

## Input Data
	
## Expected Output

 <code>{"VID":4,"UUID":"","Username":"EterniaLogic","Description":"Our new offsite storage server will hold 160TB (expandable to 264TB)!","Longitude":"33º34'57.32\" N","Latitude":"101º52'42.96\" W","Location":"Lubbock, TX","Likes":0,"Dislikes":0,"Like":false,"Dislike":false,"Sponsored":false,"Date":"2016-02-17 05:46:33"}</code>
 
 ## Test Curl
 
<code>curl -X GET -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: b2a27f56-28e2-a5bb-9acb-52d2bb36c666" "http://api.site.com/api/v1/videos/4"</code>