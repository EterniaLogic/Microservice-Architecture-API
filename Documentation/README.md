# Backend

This documents the server setup and endpoints.

The API is [REST API](http://en.wikipedia.org/wiki/Representational_State_Transfer "RESTful")
and uses [OAuth](http://oauth.net/ "OAuth") 1.0a for user authentication purposes.
Currently, return format for all endpoints is [JSON](http://json.org/ "JSON").

***


***

## Setup

- **[Server Initialization](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/docs/server_initialization.md)**
- **[Configuration](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/docs/configuration.md)**
- **[GNATsd Config](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/docs/gnatsdconfig.md)**

## Endpoints


#### Authentication

- **[<code>POST</code> api/v1/auth/user](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/auth/POST_user.md)**
- **[<code>POST</code> api/v1/auth/user/login](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/auth/GET_user_login.md)**
- **[<code>DELETE</code> api/v1/auth/user/login](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/auth/DELETE_user_login.md)**
- **[<code>GET</code> api/v1/auth/user/name/:id](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/auth/GET_user_name.md)**
- **[<code>PUT</code> api/v1/auth/oauth](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/auth/PUT_user_oauth.md)**


#### Comments

- **[<code>POST</code> api/v1/comments](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/comments/POST_comments.md)**
- **[<code>GET</code> api/v1/comments/:vid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/comments/GET_comments.md)**


#### Feedback

- **[<code>POST</code> api/v1/feedback](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/feedback/POST_feedback.md)**


#### Feeds

- **[<code>GET</code> api/v1/feeds/recent/:from/:to](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/feeds/GET_recent.md)**
- **[<code>GET</code> api/v1/feeds/follow](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/feeds/GET_follow.md)**
- **[<code>POST</code> api/v1/feeds/follow/:uid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/feeds/POST_follow.md)**
- **[<code>DELETE</code> api/v1/feeds/follow/:uid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/feeds/DELETE_follow.md)**

#### Profiles

Private:
- **[<code>GET</code> api/v1/profiles/email](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_email.md)**
- **[<code>GET</code> api/v1/profiles/firstname](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_firstname.md)**
- **[<code>GET</code> api/v1/profiles/lastname](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_lastname.md)**
- **[<code>GET</code> api/v1/profiles/gender](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_gender.md)**
- **[<code>GET</code> api/v1/profiles/city](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_city.md)**
- **[<code>GET</code> api/v1/profiles/state](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_state.md)**
- **[<code>GET</code> api/v1/profiles/country](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_country.md)**
- **[<code>PUT</code> api/v1/profiles/email](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_email.md)**
- **[<code>PUT</code> api/v1/profiles/firstname](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_firstname.md)**
- **[<code>PUT</code> api/v1/profiles/lastname](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_lastname.md)**
- **[<code>PUT</code> api/v1/profiles/gender](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_gender.md)**
- **[<code>PUT</code> api/v1/profiles/city](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_city.md)**
- **[<code>PUT</code> api/v1/profiles/state](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_state.md)**
- **[<code>PUT</code> api/v1/profiles/country](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_country.md)**


Public:
- **[<code>GET</code> api/v1/profiles/picture](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_picture.md)**
- **[<code>GET</code> api/v1/profiles/bio](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_bio.md)**
- **[<code>GET</code> api/v1/profiles/facebook](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_facebook.md)**
- **[<code>GET</code> api/v1/profiles/twitter](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_twitter.md)**
- **[<code>GET</code> api/v1/profiles/skype](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_skype.md)**
- **[<code>GET</code> api/v1/profiles/whatsapp](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_whatsapp.md)**
- **[<code>GET</code> api/v1/profiles/snapchat](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_snapchat.md)**
- **[<code>GET</code> api/v1/profiles/instagram](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_instagram.md)**
- **[<code>GET</code> api/v1/profiles/kik](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_kik.md)**
- **[<code>GET</code> api/v1/profiles/website](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_website.md)**
- **[<code>GET</code> api/v1/profiles/gear](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_gear.md)**
- **[<code>GET</code> api/v1/profiles/birthday](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/GET_birthday.md)**
- **[<code>PUT</code> api/v1/profiles/picture](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_picture.md)**
- **[<code>PUT</code> api/v1/profiles/bio](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_bio.md)**
- **[<code>PUT</code> api/v1/profiles/facebook](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_facebook.md)**
- **[<code>PUT</code> api/v1/profiles/twitter](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_twitter.md)**
- **[<code>PUT</code> api/v1/profiles/skype](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_skype.md)**
- **[<code>PUT</code> api/v1/profiles/whatsapp](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_whatsapp.md)**
- **[<code>PUT</code> api/v1/profiles/snapchat](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_snapchat.md)**
- **[<code>PUT</code> api/v1/profiles/instagram](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_instagram.md)**
- **[<code>PUT</code> api/v1/profiles/kik](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_kik.md)**
- **[<code>PUT</code> api/v1/profiles/website](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_website.md)**
- **[<code>PUT</code> api/v1/profiles/gear](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_gear.md)**
- **[<code>PUT</code> api/v1/profiles/birthday](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/profiles/PUT_birthday.md)**

#### Search

- **[<code>POST</code> api/v1/search](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/search/GET_search.md)**
- **[<code>GET</code> api/v1/search/tags/:vid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/search/GET_search_tags.md)**
- **[<code>PUT</code> api/v1/search/tags/:vid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/search/PUT_search_tags.md)**

#### Videos

- **[<code>POST</code> api/v1/videos](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/POST.md)**
- **[<code>GET</code> api/v1/videos/:vid](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/GET.md)**
- **[<code>PUT</code> api/v1/videos/desription](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/PUT_description.md)**
- **[<code>GET</code> api/v1/videos/desription](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/GET_description.md)**
- **[<code>GET</code> api/v1/videos/top/:num/:tonum](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/GET_top.md)**
- **[<code>GET</code> api/v1/videos/recent/:num/:tonum](https://github.com/EterniaLogic/Microservice-Architecture-APIblob/master/Documentation/endpoints/videos/GET_recent.md)**
