# <code>POST</code> /api/v1/auth/user

## Token Required
	none

## Input Data

#### {"Username":"JohnDoe","UserPass":"test","Email":"johndoe@gmail.com"}


## Expected Output

#### {"Type":"Success", "Message":"User created"}
#### {"Type":"Failure", "Message":"User exists"}

## Test Curl

<code>curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: c070dbeb-844f-b072-9bcb-d560c2ea5cd6" -d '{"Username":"Johndoe","UserPass":"test"}' "http://api.site.com/api/v1/auth/login"</code>