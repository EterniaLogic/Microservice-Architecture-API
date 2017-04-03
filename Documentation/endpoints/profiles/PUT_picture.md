# <code>PUT</code> /api/v1/profiles/picture

## Token Required
	none

## Input Data

<code>{"Picture": "http://ichef-1.bbci.co.uk/news/976/media/images/83351000/jpg/_83351965_explorer273lincolnshirewoldssouthpicturebynicholassilkstone.jpg"}</code>

## Expected Output

<code>{"Type":"Success", "Message":"Set Picture to http://ichef-1.bbci.co.uk/news/976/media/images/83351000/jpg/_83351965_explorer273lincolnshirewoldssouthpicturebynicholassilkstone.jpg"}</code>

## Test Curl

<code>curl -X PUT -H "Content-Type: application/json" -H "GSTX: 99E9B02F0D456210A5803201C374E3816F850E7B1ED445823F47831A" -H "Cache-Control: no-cache" -H "Postman-Token: ac5dc625-e883-b0cd-82d6-fe3338b6b687" -d '{"Picture": "http://ichef-1.bbci.co.uk/news/976/media/images/83351000/jpg/_83351965_explorer273lincolnshirewoldssouthpicturebynicholassilkstone.jpg"}' "http://api.site.com/api/v1/profiles/picture"</code>