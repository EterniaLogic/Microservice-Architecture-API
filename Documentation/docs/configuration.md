# Server Configuration

## Base File for configuration

``` json
{
	"ServerType": "Auth",
	"_Types": "Admin,Auth,Comment,Feedback,Feeds,Profile,Search,Video",
	
	"DBServer": "127.0.0.1:3306",
	"DB": "go_serv",
	"DBTable": "account",
	"DBLoginTable": "account",
	"DBUser": "",
	"DBPass": "",
	
	"NATSCluster": ["nats://a3cqExAdz1om:GOXxGx3HamsCZtu2LlEW@sync1i:8080"],
	"NATSUser": "a3cqEAdz1om",
	"NATSPass": "GOXxG3HamsCZtu2LlEW",
	
	"RESTfulPort": "6100",
	
	"_Comment":"Auth Server Only",
	"LoginTTL": 3600,
	"MailURL":"",
	"MailKey":"",
	"DoEmail":true
}
```


# Parameters

#### ServerType

This will pick a specific server type to load in code.


## Database

#### DBServer

The database server, usually 127.0.0.1:3306

#### DB

DB Schema to use. This is set to `go_serv` for the Server Initialization settings.

#### DBTable

The base name for the database, this is normally the lower case of `ServerType`.

#### DBLoginTable

The DBLoginTable is the table that helps track what users are logged in between servers. This is typically `account` on the auth server and `login` on other servers.

#### DBUser

It is bad to use the root user for the DB user, so it is better to create a user, such as `goserv` without permissions to drop tables and grant permissions.

#### DBPass

Password for the `DBUser`


## NATS

#### NATSCluster

The `NATSCluster` array allows for multiple servers to be clustered together.

#### NATSUser

The `NATSUser` option is set in the GNATSD config file for extra protection.

#### NATSPass

The `NATSPass` is along with `NATSUser` in the GNATSD config.


## TTL & Email

#### LoginTTL

The `LoginTTL` parameter is the number of seconds that a user may be logged in before their token expires

#### MailURL, MailKey

`MailURL` and `MailKey` are used for `Mailgun`.

#### DoEmail

`DoEmail` will prevent a server from sending emails if set to false, else true