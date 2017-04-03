# <code>PUT</code> /api/v1/profiles/email

## Token Required
	Header "gstx"

## Input Data

<code>{"email": "eternialogic@gmail.com"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Email to eternialogic@gmail.com"}</code>
 
 ## Testing Curl
 
 <code>curl -X PUT -H "Content-Type: application/json" -H "gstx: 4DA534852BE37FE76C1AC533CF7C9FE8B93023FDA2D959B5B3E69C12" -d '{"Email":"eternialogic@gmail.com"}' "http://10.132.19.234:6105/api/v1/profiles/email"</code>