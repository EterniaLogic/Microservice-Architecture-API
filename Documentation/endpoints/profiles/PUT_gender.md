# <code>PUT</code> /api/v1/profiles/gender

## Token Required
	Header "gstx"

## Input Data

<code>{"Gender": "Male"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Gender to Male"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: fb5fe445-f000-12c8-00c0-e0782fc9ea83" -d '{"Gender":"Male"}' "http://api.site.com/api/v1/profiles/gender"</code>