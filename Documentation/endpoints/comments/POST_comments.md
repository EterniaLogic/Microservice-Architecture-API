# <code>PUT</code> /api/v1/comments

## Token Required
	Header "gstx"

## Input Data

<code>{"VID":4,"Comment":"I liked this! \"Too awesome!\""}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Comment created"}</code>

<code>{"Type":"Failure", "Message":"Cannot post a comment"}</code>

<code>{"Type":"Failure", "Message":"Foul language"}</code>

 ## Test Curl
 
<code></code>