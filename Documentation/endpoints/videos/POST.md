# <code>POST</code> /api/v1/videos

## Token Required
	Header "gstx"

## Input Data

<code>{"VidLink":"https://www.youtube.com/watch?v=EDnAf2w2v-Y",
	"UUID":"13a6dcf3d52311e58aee0401a8d8aa01",
	"Username":"EterniaLogic",
	"Description":"Our new offsite storage server will hold 160TB (expandable to 264TB)!",
	"Longitude":"33º34'57.32\" N", "Latitude":"101º52'42.96\" W","Location":"Lubbock, TX"}</code>
			
	
## Expected Output

<code>{"Type":"Success", "Message":"Video created"}</code>

<code>{"Type":"Failure", "Message":"Video could not be created"}</code>

<code>{"Type":"Failure", "Message":"Video exists"}</code>

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: 969a0efb-4ce4-0080-5dc3-3bed07a17813" -d '{"VID":"EDnAf2w2v-Y",
	"VidLink":"https://www.youtube.com/watch?v=EDnAf2w2v-Y",
	"UUID":"13a6dcf3d52311e58aee0401a8d8aa01",
	"Username":"EterniaLogic",
	"Description":"Our new offsite storage server will hold 160TB (expandable to 264TB)!",
	"Longitude":"33º34'57.32\" N", "Latitude":"101º52'42.96\" W","Location":"Lubbock, TX"}' "http://api.site.com/api/v1/videos"</code>