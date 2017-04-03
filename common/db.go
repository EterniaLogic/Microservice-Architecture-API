package common

// User: Brent Clancy (EterniaLogic)
// Date: 12/11/2015

import (
	"log"
	"database/sql"
	"gopkg.in/gorp.v1"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // go get github.com/go-sql-driver/mysql
)

// REF:
// http://go-database-sql.org/
var db *sql.DB;
var dbmap *gorp.DbMap;

// open up a link to the database
func OpenDB(db_user string, db_pass string, db_schema string, db_connect string){    
    log.Print("Initializing database...  "+db_connect);
	LogToDBFile("DBLog.txt",true);
    
    // open the database
    dbx, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_connect+")/"+db_schema);
    db = dbx;
	
	// construct a gorp DbMap
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine:"InnoDB"}}

    
    // report an error if the variable was not even declared
    if err != nil {
            log.Fatal(err)
    }
    
    // test to make sure that the database has been connected
    err = db.Ping()
    if err != nil {
		AsyncPrintln("DB Fatal");
		log.Fatal(err)
    }else{
        AsyncPrintln("DB Opened");
    }
	
	AddTables();
    
    return
}

func OpenDBCFG(){
	OpenDB(GetConfig().DBUser,
		GetConfig().DBPass,
		GetConfig().DB,
		GetConfig().DBServer);
}

// Add tables if they do not exist based on the config
func AddTables(){
	// ServerType, DBTable, DBLoginTable
	config := GetConfig();
	ServerType := config.ServerType;
	DBTable := config.DBTable;
	DBLoginTable := config.DBLoginTable;
	
	// automatically create tables if they do not exist	
	switch(ServerType){
		case "Auth":
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`UUID` binary(16) NOT NULL,`username` varchar(50) NOT NULL,`password` blob NOT NULL,`oauth2_token` varchar(45) NOT NULL,`oauth2_website` varchar(45) NOT NULL,`oauth2_id` varchar(45) NOT NULL,`auth_token` blob NOT NULL,`auth_level` varchar(25) NOT NULL DEFAULT '',`auth_level_gen` blob NOT NULL,`join_date` datetime NOT NULL,`last_login` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',`TTL` int(11) NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='Store direct user accounts here.\nThis includes accounts that use OAuth keys generated from google, facebook.\n- Brent Clancy (EterniaLogic)';");
			break;
		case "Comment":
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`ID` int(11) NOT NULL AUTO_INCREMENT,`UUID` VARCHAR(100) NOT NULL,`VID` blob NOT NULL,`comment` varchar(45) NOT NULL,`date` datetime NOT NULL,`IP` varchar(16) NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
			break;
		case "Feedback":
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`Id` int(11) NOT NULL AUTO_INCREMENT,`UUID` binary(16) NOT NULL,`type` varchar(45) NOT NULL,`title` text NOT NULL,`action` text NOT NULL,`description` text NOT NULL,`date` datetime NOT NULL,PRIMARY KEY (`Id`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
			break;
		case "Feeds":
				//db.Exec("");
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`UUID` binary(16) NOT NULL,`followed_UUIDs` blob NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
				
			break;
		case "Profile":
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`UUID` binary(16) NOT NULL,`email` varchar(90) CHARACTER SET latin1 NOT NULL,`firstname` varchar(90) CHARACTER SET latin1 NOT NULL,`lastname` varchar(90) CHARACTER SET latin1 NOT NULL,`gender` varchar(10) NOT NULL,`picture` text NOT NULL,`bio` text CHARACTER SET latin1 NOT NULL,`city` text NOT NULL,`state` text NOT NULL,`country` text NOT NULL,`facebook` text NOT NULL,`twitter` text NOT NULL,`skype` text NOT NULL,`whatsapp` text NOT NULL,`snapchat` text NOT NULL,`instagram` text NOT NULL,`kik` text NOT NULL,`website` text CHARACTER SET latin1 NOT NULL,`gear` text CHARACTER SET latin1 NOT NULL,`birthday` date NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;");
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"_viewedvideos` (`VID` blob NOT NULL,`UUID` binary(16) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
			break;
		case "Search":
				//db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`UUID` binary(16) NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
				//db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"_tags` (`VID` INT(11) NOT NULL AUTO_INCREMENT,`tags` text NOT NULL,PRIMARY KEY (`VID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
				
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"_videos` (`VID` int(11) NOT NULL AUTO_INCREMENT,`VidLink` VARCHAR(100) NOT NULL,`UUID` binary(16) NOT NULL,`username` varchar(45) NOT NULL,`description` varchar(700) NOT NULL,`longitude` varchar(30) NOT NULL,`latitude` varchar(30) NOT NULL,`location` text NOT NULL,`date` datetime NOT NULL,`sponsored` tinyint(4) NOT NULL DEFAULT '0',`views` int(11) NOT NULL DEFAULT '0',`tags` text NOT NULL,PRIMARY KEY (`VID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='VID - Youtube video ID\ntitle - Title for the video, possibly from youtube API\nDescription - description for the video, possibly from youtube API';");
			break;
		case "Video":
				db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"_liked` (`ID` int(11) NOT NULL AUTO_INCREMENT,`UUID` binary(16) NOT NULL,`vid` int(11) NOT NULL,`islike` tinyint(1) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
			break;
	}
	
	// Videos Table
	if(ServerType == "Video" || ServerType == "Feeds" || ServerType == "Comment" || ServerType == "Search"){
		if(ServerType=="Video"){
			db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"` (`VID` int(11) NOT NULL AUTO_INCREMENT,`VidLink` VARCHAR(100) NOT NULL,`UUID` binary(16) NOT NULL,`username` varchar(45) NOT NULL,`description` varchar(700) NOT NULL,`longitude` varchar(30) NOT NULL,`latitude` varchar(30) NOT NULL,`location` text NOT NULL,`date` datetime NOT NULL,`sponsored` tinyint(4) NOT NULL DEFAULT '0',`views` int(11) NOT NULL DEFAULT '0',PRIMARY KEY (`VID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='VID - Youtube video ID\ntitle - Title for the video, possibly from youtube API\nDescription - description for the video, possibly from youtube API';");
		}else{
			db.Exec("CREATE TABLE IF NOT EXISTS `"+DBTable+"_videos` (`VID` int(11) NOT NULL AUTO_INCREMENT,`VidLink` VARCHAR(100) NOT NULL,`UUID` binary(16) NOT NULL,`username` varchar(45) NOT NULL,`description` varchar(700) NOT NULL,`longitude` varchar(30) NOT NULL,`latitude` varchar(30) NOT NULL,`location` text NOT NULL,`date` datetime NOT NULL,`sponsored` tinyint(4) NOT NULL DEFAULT '0',`views` int(11) NOT NULL DEFAULT '0',PRIMARY KEY (`VID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='VID - Youtube video ID\ntitle - Title for the video, possibly from youtube API\nDescription - description for the video, possibly from youtube API';");
		}
	}
	
	// Login Table
	if(ServerType == "Comment" || ServerType == "Feedback" || ServerType == "Feeds" || ServerType == "Profile" || ServerType == "Video" || (ServerType == "Auth" && DBTable!=DBLoginTable)){
		// Add the login table
		db.Exec("CREATE TABLE IF NOT EXISTS `"+DBLoginTable+"` (`UUID` binary(16) NOT NULL,`username` VARCHAR(90) NOT NULL,`auth_token` blob NOT NULL,`auth_level_gen` blob NOT NULL,PRIMARY KEY (`UUID`)) ENGINE=InnoDB DEFAULT CHARSET=latin1;");
	}
}

// get a column's data by either id, or token if is not empty
// returns:
//      id - UUID
//      column - list of columns separated by ,
//      output - column data for the user
//      err - true if an error has occured
func GetRowColumn(table string, id string, column string, idcol string)(output string, err bool){
	strA := "";
	strB := "";
	
	if(idcol == "UUID"){
		strA = "UNHEX(";
		strB = ");";
	}
	
	errx := GetDB().QueryRow("SELECT "+column+" FROM `"+table+"` WHERE `"+idcol+"`="+strA+"?"+strB+";",id).Scan(&output);
	
	queryq := "SELECT "+column+" FROM `"+table+"` WHERE `"+idcol+"`="+strA+"'"+id+"'"+strB+"');";
	DBAsyncPrintln(queryq);
	
	// check the size
	if(errx != nil){
		err=true;
		//common.AsyncPrintln("GetRowColumnWhere: ERROR=",errx);
		DBAsyncPrintln("GetRowColumnWhere: ERROR="+fmt.Sprintf("%s", errx));
	}
    return;
}

func GetRowColumnNE(table string, id string, column string,idcol string)(string){
	out,_ := GetRowColumn(table,id,column,idcol);
    return out;
}

func GetRowColumnWhere(table string,wherecolumn string,where string,column string) (output string){
	errx := GetDB().QueryRow("SELECT "+column+" FROM `"+table+"` WHERE `"+wherecolumn+"`=?;",where).Scan(&output);
	
	queryq := "SELECT "+column+" FROM `"+table+"` WHERE `"+wherecolumn+"`='"+where+"';";
	DBAsyncPrintln(queryq);
	
	// check the size
	if(errx != nil){
		//common.AsyncPrintln("GetRowColumnWhere: ERROR=",errx);
		DBAsyncPrintln("GetRowColumnWhere: ERROR="+fmt.Sprintf("%s", errx));
	}
	return;
}

// ID is a HEX value of a binary number
func RowExists(ID string, COL string, PREFIX string, requiresUNHEX bool)(bool){
	var IDx string;
	strA := "";
    strD := "";
	strE := "";
	if(requiresUNHEX){
		strA = "HEX("
		strD = "UNHEX(";
		strE = ")";
	}
	
	errx:=GetDB().QueryRow("SELECT "+strA+COL+strE+" FROM `"+GetConfig().DBTable+PREFIX+"` WHERE `"+COL+"`="+strD+"?"+strE+" LIMIT 1;",ID).Scan(&IDx);
    
	queryq := "SELECT "+strA+COL+strE+" FROM `"+GetConfig().DBTable+PREFIX+"` WHERE `"+COL+"`="+strD+ID+strE+" LIMIT 1;";
	DBAsyncPrintln(queryq);
	
    // check the size
    if(errx != sql.ErrNoRows || ID == IDx){
		DBAsyncPrintln("RowExists: ERROR="+fmt.Sprintf("%s", errx));
		return true;
	}else{
		return false;
	}
}

// ID is a HEX value of a binary number
func RowExists2(ID string, COL string, ID2 string, COL2 string, PREFIX string)(bool){
	var IDx string;
	strA := "";
    strB := "";
	strC := "";
    strE := "";
	strF := "";
	if(COL=="UUID"){
		strA = "HEX("
		strB = "UNHEX(";
		strC = ")";
	}
	if(COL2=="UUID"){
		strE = "UNHEX(";
		strF = ")";
	}
	
    errx:=GetDB().QueryRow("SELECT "+strA+COL+strC+" FROM "+GetConfig().DBTable+PREFIX+" WHERE `"+COL+"`="+strB+"?"+strC+" AND `"+COL2+"`="+strE+"?"+strF+" LIMIT 1;",ID,ID2).Scan(&IDx);
    
    // check the size
    if(errx != sql.ErrNoRows || ID == IDx){
		return true;
	}else{
		return false;
	}
}

// returns the database reference
func GetDB()(*sql.DB){
	return db;
}

func GetDBGorp()(*gorp.DbMap){
	return dbmap;
}

// close the database
func CloseDB(){
	CloseDBLogToFile();
    if(db != nil){
        db.Close();
    }
}