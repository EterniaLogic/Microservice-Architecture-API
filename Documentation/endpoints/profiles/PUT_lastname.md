# <code>PUT</code> /api/v1/profiles/lastname

## Token Required
	Header "gstx"

## Input Data

<code>{"Lastname": "eternialogic@gmail.com"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Lastname to Logic"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 49a12b69-b15f-bffa-589c-98a64e1e5479" -d '{"Firstname":"Logic"}' "http://api.site.com/api/v1/profiles/firstname"</code>