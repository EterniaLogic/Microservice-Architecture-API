package common

import "testing"
import "log"
import "time"

// types listed below are from nats_types.go

// test NATS
func TestNATSServer(t *testing.T){
	StartNATSClient();
	
	// an ID is a UUID without dashes
	auth:=ConfirmUserLevel("CC", "Administrator", "CCZCXX");
	
	if(auth == true){
		common.AsyncPrintln("TestNATSServer Error: unknown user somehow got true authorization")
	}
}



func TestUserManagerLogin(t *testing.T){
	InitUserManager();
	
	ListenToLogin(func (ID string){
		common.AsyncPrintln("Login Listener:",ID);
	});
	
	ListenToLogout(func (ID string){
		common.AsyncPrintln("Logout Listener:",ID);
	});
	
	ttl := time.Now().Unix()+5;
	common.AsyncPrintln("TEST ttl:",ttl);
	loginmsg := AuthLogin{ID:"EE0D9E54A92611E5972C04018E7C6601",AuthLevel:"User",HashedAuthLevel:"ttttttt", TTL:ttl};
	GetNATSClient().Publish("Auth.Login",loginmsg);
	loginmsg = AuthLogin{ID:"EE0D9E54A92611E5972C04018E7C6601",AuthLevel:"User",HashedAuthLevel:"ttttttt", TTL:ttl};
	GetNATSClient().Publish("Auth.Login",loginmsg);
	time.Sleep(10*time.Second);
}


func TestCloseNats(t *testing.T){
	CloseNATSClient();
}