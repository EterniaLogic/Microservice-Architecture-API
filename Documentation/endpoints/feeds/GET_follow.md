# <code>GET</code> /api/v1/feeds/follow

## Token Required
	Header "gstx"

## Input Data

## Expected Output

 <code>{"Number":1,"Following":["13A6DCF3D52311E58AEE0401A8D8AA01"]}</code>
 
 <code>{"Type":"Failure","Message":"Unable to get followed users}</code>
 
## Test Curl
 
 <code>curl -X GET -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: f84c848c-77a9-bf7e-4eab-29bf9a1f13a0" "http://api.site.com/api/v1/feeds/follow"</code>