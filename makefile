all:
	go install client_main/client_main.go	
	go install server_main/server_main.go
	go build -o plugins -buildmode=plugin test_plugin/*	
c:	
	go install client_main/client_main.go
s:
	go install server_main/server_main.go
plug:  
	#$(shell rm ./plugins/*)
	go build -o plugins -buildmode=plugin test_plugin/*
	bash make_plug.sh
p:
	bash make_proto.sh
