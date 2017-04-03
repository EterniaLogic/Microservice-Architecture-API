# <code>PUT</code> /api/v1/auth/user/oauth/

## Token Required
	Header "gstx"

## Input Data

#### {"OAuth2Token":"8eff1b488da7fe3426f9ecaf8de1ba54"}

## Expected Output

#### {"Type":"Success", "Message":"Successfully placed OAuth2 Token."}
#### {"Type":"Failure", "Message":"ID does not exist"}