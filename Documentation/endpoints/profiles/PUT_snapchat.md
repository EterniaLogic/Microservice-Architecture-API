# <code>PUT</code> /api/v1/profiles/snapchat

## Token Required
	none

## Input Data

<code>{"Snapchat": "eternialogic"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Snapchat to eternialogic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 0686983f-c6dd-8ad8-e6b8-47a710043f0a" -d '{"Snapchat": "eternialogic"}' "http://api.site.com/api/v1/profiles/snapchat"</code>