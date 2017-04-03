package common

// User: Brent Clancy (EterniaLogic)
// Date: 1/8/2016

// Sent to the Auth Server from any server to determine the level
type AuthVerifyLevelMsg struct{
	ID string;
	Wanted string;	// "User","Administrator"
	Hashed string;  // Value given from AuthLogin
}

// return message from Verify Level
type AuthVerifyLevelRetMsg struct{
	ID string;
	Authorized bool;
}

// pushed out from the auth server when somebody has logged in
// on the "Auth.Login" channel on NATS
type AuthLogin struct{
	ID string; // UUID
	Username string; // EterniaLogic or whatever
	Token string; // token for the user
	HashedAuthLevel string;
}

type Admin struct{
	UUID string;		// UUID HEX for the admin
	Username string;
	IP string;			// IP address for the admin
	IDToken string;
	IsAdmin string;		// Verified admin
	UserLevel string;	// plaintext userlevel
	UserLevelGen string;// Generated HEX userlevel
};

type VerifyAdmin struct{
	IP string;
	Username string;
	Password string; // pre-hashed hashed password
};

// used for JSON output from AuthLogin
type Authentication struct{
    Type string;
    ID string;
    Token string;
}

// UserLevel JSON used for higher priviledges
type UserLevel struct{
    Type string;
    UserLevel string;
    HashedUserLevel string;
}

// UserProfile JSON output
type UserProfile struct{
    Type string;
    RealName string;
    LastLogin string;
    UserLevel string;
}

// used for basic JSON output messages
type MsgData struct{
    Type string;
    Message string;
}

// Generic JSON output
type ValueX struct{
    Type string;
    Value string;
}

type SingleValue struct{
	Value string;
};

type VideoData struct{
	VID int64;
	UUID string;
	VidLink string;
	Username string;
	Description string;
	Longitude string;
	Latitude string;
	Location string;
	Likes int;
	Dislikes int;
	Like bool;
	Dislike bool;
	Sponsored bool;
	Date string;
	Views int;
}

type VideoView struct{
	VID int64;
	UUID string;
};

type CommentX struct{
	UUID string;
	Comment string;
	Date string;
};