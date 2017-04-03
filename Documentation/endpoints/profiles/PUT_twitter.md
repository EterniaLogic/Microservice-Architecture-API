# <code>PUT</code> /api/v1/profiles/twitter

## Token Required
	none

## Input Data

<code>{"Twitter": "eternialogic"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Twitter to eternialogic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 7f8a7ff0-b69f-8fee-f27d-5e818f5100d2" -d '{"Twitter": "eternialogic"}' "http://api.site.com/api/v1/profiles/twitter"</code>