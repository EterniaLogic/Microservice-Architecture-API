# <code>POST</code> /api/v1/feedback

## Token Required
	Header "gstx"

## Input Data

<code>{"Type":"Bug", "Title":"Why no workie!","Action":"Watching video", "Description":"Videos all borked up!"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Feedback added"}</code>

<code>{"Type":"Failure", "Message":"Cannot post feedback... bug for reporting bugs?"}</code>

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: ae83217d-de81-bf77-3786-0b7363e2b7e2" -d '{"Type":"Bug", "Title":"Why no workie!","Action":"Watching video", "Description":"Videos all borked up!"}' "http://api.site.com/api/v1/feedback"</code>