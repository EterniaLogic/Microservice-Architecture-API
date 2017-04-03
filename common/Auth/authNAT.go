package Auth

import "../../common"

func InitSubscribeNAT(){
	common.GetNATSClientJSON().Subscribe("Auth.VerifyLevel", func(subj string, reply string, msg *common.AuthVerifyLevelMsg) {
		go func(){
			_, err := AuthVerifyUserLevel(msg.ID, msg.Wanted, msg.Hashed);
			Msg := common.AuthVerifyLevelRetMsg{ID:msg.ID,Authorized:(err==false)};
			common.GetNATSClientJSON().Publish(reply, Msg);
		}();
	});
	
	common.GetNATSClientJSON().Subscribe("Auth.Admin.Login", func(subj string, reply string, msg *common.VerifyAdmin) {
		go func(){
			// Admin.UUID, Admin.Password
			admin := common.Admin{UUID:"",Username:msg.Username,IP:msg.IP};
			
			id,idtoken,_,err := AuthUser(msg.Username, msg.Password, true);
			admin.UUID = id;
			
			if(idtoken != "" && !err){
				admin.IDToken = idtoken;
				admin.IsAdmin = common.PreSum("Bdkijw");
				
				var UUID, auth_level, auth_level_gen string;
				common.GetDB().QueryRow("SELECT HEX(UUID),auth_level,HEX(auth_level_gen) FROM `"+common.GetConfig().DBTable+"` WHERE username=?;",msg.Username).Scan(&UUID, &auth_level, &auth_level_gen);
				
				admin.UUID = UUID;
				admin.UserLevel = auth_level;
				admin.UserLevelGen = auth_level_gen;
				
				common.GetNATSClientJSON().Publish(reply, admin);
			}
		}();
	});
}



/*
	type Admin struct{
		UUID string;		// UUID HEX for the admin
		Username string;
		IP string;			// IP address for the admin
		IDToken string;
		IsAdmin string;		// Verified admin
		UserLevel string;	// plaintext userlevel
		UserLevelGen string;// Generated HEX userlevel
	};
*/
