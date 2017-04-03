# <code>PUT</code> /api/v1/profiles/website

## Token Required
	none

## Input Data

<code>{"Website": "https://eternialogic.com/"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Website to https://eternialogic.com/"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 82c933fa-23cb-54bd-ed13-0ab113429884" -d '{"Website": "https://eternialogic.com/"}' "http://api.site.com/api/v1/profiles/website"</code>