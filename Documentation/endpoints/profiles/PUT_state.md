# <code>PUT</code> /api/v1/profiles/state

## Token Required
	Header "gstx"

## Input Data

<code>{"State": "Texas"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set State to Texas"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 8db9979a-98c5-a924-e6b1-7d643538c22a" -d '{"State": "Texas"}' "http://api.site.com/api/v1/profiles/state"</code>