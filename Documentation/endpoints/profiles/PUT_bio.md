# <code>PUT</code> /api/v1/profiles/bio

## Token Required
	none

## Input Data

<code>{"Bio": "Bio goes here!"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Bio to Stuff goes here!"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 08bd8ae6-a6d7-bd61-7394-2416ee0edf23" -d '{"Bio":"Stuff goes here!"}' "http://api.site.com/api/v1/profiles/bio"</code> 