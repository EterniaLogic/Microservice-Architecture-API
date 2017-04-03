# <code>DELETE</code> /api/v1/feeds/follow/:uid

## Token Required
	Header "gstx"

## Input Data

## Expected Output

<code>{"Type":"Success", "Message":"Unfollowed user"}</code>

<code>{"Type":"Failure", "Message":"ID does not exist"}</code>

## Test Curl

<code>curl -X GET -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 4911b8d1-f24c-f780-613f-fe5b4244b3fb" "http://api.site.com/api/v1/feeds/follow"</code>