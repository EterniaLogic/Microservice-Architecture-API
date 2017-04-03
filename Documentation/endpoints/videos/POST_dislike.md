# <code>POST</code> /api/v1/videos/dislike/:vid

## Token Required
	Header "gstx"

## Input Data
			
	
## Expected Output

<code>{"Type":"Success", "Message":"Disliked Video"}</code>

<code>{"Type":"Success", "Message":"Undisliked Video"}</code>

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "Gstx: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 22225409-4cdd-1db0-0df7-c71fa793751f" "http://api.site.com/api/v1/videos/dislike/4"</code>