# <code>PUT</code> /api/v1/profiles/country

## Token Required
	Header "gstx"

## Input Data

<code>{"Country": "United States"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Country to United States"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 4f825154-b193-f58a-2221-2aeda3a32dd9" -d '{"Country":"United States"}' "http://api.site.com/api/v1/profiles/country"</code>