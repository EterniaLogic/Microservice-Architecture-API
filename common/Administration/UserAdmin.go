package Administration

import "../../common"

// add a new user
func (a Admin) AddNewUser(username string, email string)(bool){
	// Verify that this is indeed an admin
	if(a.VerifyUser("Administrator")){
		common.AsyncPrintln("[Admin] Adding new User: "+username+", "+email);
	}else{
		return false;
	}
	
	return true;
}

// remove a user
func (a Admin) RemoveUser(UUID string)(bool){
	// Verify that this is indeed an admin
	if(a.VerifyUser("Administrator")){
		common.AsyncPrintln("[Admin] Removing User: "+UUID);
	}else{
		return false;
	}
	
	return true;
}

// modify a user's level
func (a Admin) ModifyUserLevel(UUID string, Level string)(bool){
	// Verify that this is indeed an admin
	if(a.VerifyUser("Administrator")){
		common.AsyncPrintln("[Admin] Modify User (Set Level `"+Level+"`): "+UUID);
	}else{
		return false;
	}
	
	return true;
}