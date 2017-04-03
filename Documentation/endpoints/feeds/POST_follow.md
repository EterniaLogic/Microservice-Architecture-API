# <code>POST</code> /api/v1/feeds/follow/:uid

## Token Required
	Header "gstx"

## Input Data

## Expected Output

<code>{"Type":"Success", "Message":"Followed user"}</code>

<code>{"Type":"Failure", "Message":"ID does not exist"}</code>

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 883d7134-be9c-0e2c-5c26-baa34eff561d" -d '' "http://api.site.com/api/v1/feeds/follow/13A6DCF3D52311E58AEE0401A8D8AA01"</code>