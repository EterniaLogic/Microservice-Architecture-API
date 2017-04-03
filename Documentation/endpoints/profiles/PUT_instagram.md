# <code>PUT</code> /api/v1/profiles/instagram

## Token Required
	none

## Input Data

<code>{"Instagram": "http://instagram.com"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Instagram to http://instagram.com"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: ec75798f-38df-eb43-413d-4aeae1b57199" -d '{"Instagram": "http://instagram.com"}' "http://api.site.com/api/v1/profiles/instagram"</code>