# <code>PUT</code> /api/v1/profiles/city

## Token Required
	Header "gstx"

## Input Data

<code>{"City": "Manhattan"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set City to Lubbock"}</code>
 
 ## Testing Curl
 
 <code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 043eea22-a7c6-bb7d-5d79-252ef0492866" -d '{"City":"Lubbock"}' "http://api.site.com/api/v1/profiles/city"</code>