package test

/*import "testing"
import "database/sql"
import "log"
import "../common"
import "../common/Auth"
import "strings"

// by opening the DB multiple times, the DB link is also tested.
// These tests are all assumed to work separately, but some require the
// test user to be alive.
func TestOpenDB(t *testing.T){
	common.SetConfig("conf.json");
	common.OpenDBCFG();
	common.StartNATSClientCFG();
	common.InitMailer();
}

func TestAuthRegister(t *testing.T){
	// first, check if the user exists in the system
	var username string;
    errx:=common.GetDB().QueryRow("SELECT username FROM account WHERE username=?;","testcase").Scan(&username);
    
    // if the user exists, remove it:
    if(errx != sql.ErrNoRows || "testcase" == username){
		//log.Println("TestAuthRegister: Removing test user first");
		common.GetDB().Exec("DELETE FROM account WHERE username=?","testcase");
	}

    // create a new user through auth
	//log.Println("TestAuthRegister: Adding new user");
	msg,err := Auth.AuthRegister("testcase", "testpwd", "eternialogic@gmail.com");
	
	if(err){
		log.Println("TestAuthRegister: Failed to add a new user!  MSG: "+msg);
		t.Fail();
	}
	
	var UUID string;
    common.GetDB().QueryRow("SELECT HEX(UUID) FROM account WHERE username=?;","testcase").Scan(&UUID);
	// set user to authorized USER
	auth := common.GenAuthLevel("User", UUID);
	common.GetDB().Exec("UPDATE account SET auth_level=?, auth_level_gen=UNHEX(?) WHERE `username`='testcase';","User",auth);
}

func TestAuth(t *testing.T){
	// authorize the testcase user!
	id, idtoken, msg, err := Auth.AuthUser("testcase", "testpwd", false);
	
	// test for errors
	if(err){
		log.Println("TestAuth: FAIL: "+msg);
		t.Fail();
	}else{
		//log.Println("TestAuth: idtoken = "+idtoken);
	}
	
	// test the outputs:
	if(idtoken == ""){
		log.Println("TestAuth: FAIL: idtoken is not defined");
		t.Fail();
	}
	
	if(id == "0"){
		log.Println("TestAuth: FAIL: id is 0, it should be higher!");
		t.Fail();
	}

	// test that the idtoken is placed in the table
	var auth_token string;
    common.GetDB().QueryRow("SELECT HEX(auth_token) FROM account WHERE username=?;","testcase").Scan(&auth_token);
    
    // if the user exists, remove it:
    if(auth_token != strings.ToUpper(idtoken)){
		log.Println("TestAuth: FAIL: idtoken is not placed in the table!");
		log.Println(auth_token,"!=",strings.ToUpper(idtoken))
		t.Fail();
	}
}

func BenchmarkAuth(b *testing.B){
	Auth.AuthUser("testcase", "testpwd", false);
}


func TestOAuth(t *testing.T){
}

func TestAuthGetColumn(t *testing.T){	
	// reauthenticate first to get the user id
	id,_,_,_ := Auth.AuthUser("testcase", "testpwd", false);
	
	//log.Println("TestAuthGetColumn: ID="+id,idtoken,msg,err);
	
	// now pull the username with that ID
	username,erry := common.GetRowColumn("account",id,"`username`");
	
	if(erry){
		log.Println("TestAuthGetColumn: FAIL USER ID not found");
		t.Fail();
	}
	
	if(username != "testcase"){
		log.Println("TestAuthGetColumn: FAIL output != input");
		t.Fail();
	}
}



func BenchmarkOAuth(b *testing.B){

}

func TestVerifyAuthLevel(t *testing.T){
	// reauthenticate first to get the user id
	id,_,msg,err := Auth.AuthUser("testcase", "testpwd", false);
	
	
	authgen,_ := common.GetRowColumn("account",id,"HEX(`auth_level_gen`)");
	
	// first test for a failure mode, a user wants admin priviledges
	msg,err = Auth.AuthVerifyUserLevel(id, "Administrator", authgen);
	
	if(!err){
		log.Println("TestVerifyAuthLevel: User=Admin? FAIL",msg);
		t.Fail();
	}
	
	// next, test if a user can do a user thing (shouldnt happen, but just to test)
	msg,err = Auth.AuthVerifyUserLevel(id, "User", authgen);
	
	if(err){
		log.Println("TestVerifyAuthLevel: User!=User? FAIL",msg);
		t.Fail();
	}
	
	
	// finally, test if an admin can do admin things
	// set admin mode:
	auth := common.GenAuthLevel("Administrator", id);
	common.GetDB().Exec("UPDATE account SET auth_level=?, auth_level_gen=UNHEX(?) WHERE `UUID`=UNHEX(?);","Administrator",auth,id);
	
	// do Admin=Admin?
	msg,err = Auth.AuthVerifyUserLevel(id, "Administrator", authgen);
	
	if(err){
		log.Println("TestVerifyAuthLevel: Admin!=Admin? FAIL",msg);
		t.Fail();
	}
}

// user logout test
func TestDeAuth(t *testing.T){
	id,_,_,_ := Auth.AuthUser("testcase", "testpwd", false);
	Auth.AuthLogout(id);
}

// Remove the test user in case it exists
func TestRemoveUser(t *testing.T){
	common.GetDB().Exec("DELETE FROM account WHERE username=?","testcase");
}

func TestCloseDB(t *testing.T){
	common.CloseDB();
}*/