#!/bin/bash
protoc protocol/protocol.proto --go_out=src/
#mv protocol.pb.go ../src/protocol 
protoc protocol/rpc.proto --go_out=plugins=grpc:src/
