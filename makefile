GOPATH := $(shell pwd)
all:
	GOPATH=$(GOPATH) go install server_main
	GOPATH=$(GOPATH) go install client_main
	go build -o plugins -buildmode=plugin src/plugins/*	
c:	
	GOPATH=$(GOPATH) go install client_main
s:
	GOPATH=$(GOPATH) go install server_main
plug:  
	#$(shell rm ./plugins/*)
	go build -o plugins -buildmode=plugin src/plugins/*
	bash make_plug.sh
p:
	bash make_proto.sh
