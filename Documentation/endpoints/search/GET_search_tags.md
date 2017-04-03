# <code>GET</code> /api/v1/search/tags/:vid

## Token Required
	none

## Input Data

## Expected Output

<code>{
  "Type": "Tags",
  "Value": "Test,woo,yay"
}</code>
 
 ## Test Curl
 
 <code>curl -X GET -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: 6e0aea01-b0b4-90e6-65d5-f0317ca8a244" "http://api.site.com/api/v1/search/tags/1"</code>