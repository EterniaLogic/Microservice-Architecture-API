# <code>GET</code> /api/v1/comments/:vid

## Token Required
	none

## Input Data


## Expected Output

<code>{"Comments":[{"UUID":"13A6DCF3D52311E58AEE0401A8D8AA01","Comment":"I liked this! 1 \"Too awesome!\"","Date":"2016-02-17 11:22:37"},{"UUID":"13A6DCF3D52311E58AEE0401A8D8AA01","Comment":"I liked this! \"Too awesome!\"","Date":"2016-02-17 11:22:36"},{"UUID":"13A6DCF3D52311E58AEE0401A8D8AA01","Comment":"I liked this! \"Too awesome!\"","Date":"2016-02-17 11:17:34"}]}</code>

## Test Curl

<code>curl -X GET -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: fd06719b-34ff-07a1-52f4-3c9134bc78a4" "http://api.site.com/api/v1/comments/4"</code>