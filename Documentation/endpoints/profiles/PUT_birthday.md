# <code>PUT</code> /api/v1/profiles/birthday

## Token Required
	none

## Input Data

<code>{"Birthday": "December 21th"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Birthday to December 21st"}</code>
 
## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 744d2d72-1bc9-0a11-00f2-d461798c795e" -d '{"Birthday":"December 21st"}' "http://api.site.com/api/v1/profiles/birthday"</code>