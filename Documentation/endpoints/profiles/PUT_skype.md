# <code>PUT</code> /api/v1/profiles/skype

## Token Required
	none

## Input Data

<code>{"Skype": "eternialogic"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Skype to eternialogic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: d6e87d9a-0c14-e58f-9b2e-6c1921d3b2d0" -d '{"Skype": "eternialogic"}' "http://api.site.com/api/v1/profiles/skype"</code>