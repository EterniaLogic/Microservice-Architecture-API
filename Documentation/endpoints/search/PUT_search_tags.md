# <code>PUT</code> /api/v1/search/tags/:vid

## Token Required
	Header "gstx"

## Input Data
	<code>{"Tags": "Test,woo,yay"}</code>
## Expected Output

 <code>{
  "Type": "Success",
  "Message": "Set Tags to Test,woo,yay"
}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 34b59512-5440-96ec-a163-dcdd06f387f6" -d '{"Tags": "Test,woo,yay"}' "http://api.site.com/api/v1/search/tags/1"</code>