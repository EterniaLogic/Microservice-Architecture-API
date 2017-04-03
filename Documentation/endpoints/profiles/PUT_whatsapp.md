# <code>PUT</code> /api/v1/profiles/whatsapp

## Token Required
	none

## Input Data

<code>{"Whatsapp": "eternialogic"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Whatsapp to eternialogic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: bcd1b35e-48ba-d612-b627-780ba0099217" -d '{"Whatsapp": "eternialogic"}' "http://api.site.com/api/v1/profiles/whatsapp"</code>