# <code>PUT</code> /api/v1/profiles/kik

## Token Required
	none

## Input Data

<code>{"Kik": "http://www.kik.com/"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Kik to http://www.kik.com/"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 9e4d5426-60ae-21ef-5219-c6141992366c" -d '{"Kik": "http://www.kik.com/"}' "http://api.site.com/api/v1/profiles/kik"</code>