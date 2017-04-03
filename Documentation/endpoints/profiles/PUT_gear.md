# <code>PUT</code> /api/v1/profiles/gear

## Token Required
	none

## Input Data

<code>{"Gear":"Go Pro v4"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Gear to Go Pro v4<br />Samsung S6 Camera"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 0375a6bc-828e-f7ac-6637-d68ec9f02027" -d '{"Gear":"Go Pro v4<br/>Samsung S6 Camera"}' "http://api.site.com/api/v1/profiles/gear"</code>