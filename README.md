# DISYS-HandIn-5

1: the program is hard-coded to work with 3 total servers/replicas, but with
an unlimited amount of clients
2: we recommend readying 3 servers and 3 clients in 6 total terminals,
thereby removing some of the time pressure of starting each individually
3: start the 3 servers with port 5000, 5001 and 5002 (it will ask for port
after the initial ”go run server.go”
4: start the three (or more) clients, they will auto connect to localhost with
ports 5000, 5001 and 5002
5: you can now bid by typing any integer, or query the status of the auction
by typing ”status”
6: after 180 sec. the auction will finish, giving the result and returning
”AUCTION OVER” to any client that tries to bid
7: to simulate a crash, simply close one of the server terminals, to reconnect
a crashed server, simply start it up again with the same port.
