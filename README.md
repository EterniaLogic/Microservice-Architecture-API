# Microservice Architecture API

Author: Brent Clancy

Copyright (C) Brent Clancy 2019

#### Documentation:

https://github.com/EterniaLogic/Microservice-Architecture-API/blob/master/Documentation/README.md


Ports:

		Admin 		port 6100
		Auth 		port 6101
		Comments 	port 6102
		Feedback 	port 6103
		Feeds 		port 6104
		Profile 	port 6105
		Search 		port 6106
		Video 		port 6107


The majority of this code is non-blocking. In that way, the server can handle multiple requests even with a single processor, or can automatically scale based on the server.
RESTFul through HTTP is used for communication with clients and GNATSD (nats.io) for inter-server communication.
