#!/bin/bash
protoc protocol/protocol.proto --go_out=./
protoc protocol/rpc.proto --go_out=plugins=grpc:./
