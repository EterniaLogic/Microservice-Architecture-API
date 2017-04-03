# <code>POST</code> /api/v1/videos/like/:vid

## Token Required
	Header "gstx"

## Input Data
			
	
## Expected Output

<code>{"Type":"Success", "Message":"Liked Video"}</code>

<code>{"Type":"Success", "Message":"Unliked Video"}</code>

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "Gstx: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: bf08a983-1cd8-a909-8304-896b720561f4" "http://api.site.com/api/v1/videos/like/4"</code>