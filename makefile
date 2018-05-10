GOPATH := $(shell pwd)
all:
	GOPATH=$(GOPATH) go install server_main
	GOPATH=$(GOPATH) go install client_main
p:

	./make_proto.sh
