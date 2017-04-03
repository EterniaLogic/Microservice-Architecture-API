Author: Brent Clancy
Date: 12/11/2015

Port: 6100   (Must be a high number for a non-root)

[TODO]
- OAuth2 support for facebook, google and twitter

[Testing]
curl -H 'Content-Type: application/json' -X POST -d '{"Username":"JohnDoe","UserPass":"test","Email":"johndoe@gmail.com"}' localhost:6101/api/v1/auth/user
curl -H 'Content-Type: application/json' -X POST -d '{"Username":"JohnDoe","UserPass":"test"}' localhost:6101/api/v1/auth/login
curl -X DELETE localhost:6100/api/v1/auth/login/EE0D9E54A92611E5972C04018E7C6601

[Testing - Burning]
while true; do curl -H 'Content-Type: application/json' -X POST -d '{"Username":"JohnDoe2”,”UserPass":"test"}' localhost:6100/api/v1/auth/login > /dev/null 2>&1 ; done

[Authorization]
The two methods for authentication through this microservice is:
- Direct
- OAuth2

Either methods will return a User ID and a token string.
- The User ID is an encrypted increment for ID in the Auth db
- The token string is given to the endpoint client to track them

OAuth2 will automatically create a user profile upon login, while 
Direct authentication will require a register step.


[Higher Permissions]
    User Level is used to prevent a user from going over their permissions.
An Administrator level must be protested through the Auth Server before
actual changes may be made to other accounts, comments, locations, ect.
This is all managed by common.usermanager.go

"Administrator" > "Moderator" > "User"

Wanted is the string name of a user level i.e: "User", "Administrator"
Comparisons are as below:
	"Administrator" = 9999
	"VideoModerator" = 101
	"CommentModerator" = 90
	"User" = 5
	"Unverified" = 1
	"Banned" = 0


================ NATS Send/Returns ================	
Auth.Register -> common.UserManager
	Sends to all servers that this user is created.
Auth.Login -> common.UserManager
	Sends to all servers that this user has logged in.
Auth.Logout -> common.UserManager
	Sends to all servers that this user has logged out.
Admin.delete.user <- common.UserManager
	Clears a user's information from this server
			
========== RESTful and JSON Send/Returns ==========

Create a user
POST /api/v1/auth/user
	{"Username":"JohnDoe","UserPass":"test","Email":"johndoe@gmail.com"}
                {"Type":"Failure", "Message","User exists"}

[Authorization]
POST /api/v1/auth/user/login
	{"Username":"JohnDoe","UserPass":"test"}
			{"Type":"Failure", "Message","Username/Password incorrect"}
			{"Type":"Authentication","ID":"EE0D9E54A92611E5972C04018E7C6601","Token":"81e84a1d425901f50abf45d4fcf6d71f6d1eee129a14f0ea17123bbd","TTL":"1450311379"}
	{"ID":"95341411595","Token":"8eff1b488da7fe3426f9ecaf8de1ba54","Website":"facebook.com","TTL":"1450311379"}
			{"Type":"Failure", "Message","Bad OAuth id/token"}
			{"Type":"Authentication","ID":"EE0D9E54A92611E5972C04018E7C6601","Token":"81e84a1d425901f50abf45d4fcf6d71f6d1eee129a14f0ea17123bbd","TTL":"1450311379"}

[Logout a user]
DELETE /api/v1/auth/user/login/{uid}
	{"Type":"Failure", "Message","User is already logged out"}
	{"Type":"Failure", "Message","ID does not exist"}
	{"Type":"Success", "Message":"User logged out"} 

[Retrieve a username]
GET /api/v1/auth/user/name/{uid}
	{"Type":"Failure", "Message","ID does not exist"}
	{"Type":"Username", "Value":"JohnDoe"}
    
[Retrieve when user last logged in]
GET /api/v1/auth/user/lastlogin/{uid}
	{"Type":"Failure", "Message","ID does not exist"}
	{"Type":"LastLogin", "Message","Dec 10, 2015 5:30 PM"}