#!/bin/bash
cd protocol
protoc --go_out=. protocol.proto
mv protocol.pb.go ../src/protocol 
