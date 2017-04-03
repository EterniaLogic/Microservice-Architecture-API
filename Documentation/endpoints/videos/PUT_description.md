# <code>PUT</code> /api/v1/videos/description/:vid

## Token Required
	Header "gstx"

## Input Data

<code>{"Description":"STUFF"}</code>
	
## Expected Output

<code>{"Type":"Success", "Message":"Set Description to STUFF"}</code> 

<code>{"Type":"Failure", "Message":"Could not modify description"}</code> 

<code>{"Type":"Failure", "Message":"Video does not exist"}</code> 

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 86d5353e-50ee-9a85-0237-30f2c09c2d51" -d '{"Description":"STUFF"}' "http://api.site.com/api/v1/videos/description/4"</code>