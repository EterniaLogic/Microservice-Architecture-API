# <code>PUT</code> /api/v1/profiles/facebook

## Token Required
	none

## Input Data

<code>{"Facebook":"https://www.facebook.com/eternialogic"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Facebook to https://www.facebook.com/eternialogic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 3cbbd9f1-4dde-ee27-dcfc-88fa629ce3e9" -d '{"Facebook":"https://www.facebook.com/eternialogic"}' "http://api.site.com/api/v1/profiles/facebook"</code>