# <code>GET</code> /api/v1/videos/description

## Token Required
	none

## Input Data

	
## Expected Output

<code>{"Type":"Description","Value":"Our new offsite storage server will hold 160TB (expandable to 264TB)!"}</code> 

<code>{"Type":"Failure", "Message":"Video does not exist"}</code> 

## Test Curl

curl -X GET -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: ae5b26f1-e581-ae24-1900-fc4d6f92b620" "http://api.site.com/api/v1/videos/description/4"