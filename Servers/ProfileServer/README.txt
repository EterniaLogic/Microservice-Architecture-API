Author: Brent Clancy
Date: 12/23/2015

Port: 6101   (Must be a high number for a non-root)


================ Testing/Debugging ================
[Testing]
wget http://localhost:6101/api/v1/profile/picture/EE0D9E54A92611E5972C04018E7C6601

[Testing - Burning the memory/cpu]
while true; do curl http://localhost:6101/api/v1/profile/picture/EE0D9E54A92611E5972C04018E7C6601 > /dev/null 2>&1 ; done
	-- results: ~8% cpu, 0.9% increase in memory after 15 minutes

http://localhost:6101/debug/pprof/heap
	Mallocs: 1318418 (hard burn)


================ NATS Send/Returns ================	
Auth.Login <- common.UserManager
	UserManager will append this information to it's internal map
Auth.Logout <- common.UserManager
	UserManager manages this.
Admin.delete.user <- common.UserManager
	Clears a user's information from this server
			
========== RESTful and JSON Send/Returns ==========
