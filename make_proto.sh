#!/bin/bash
protoc protocol/protocol.proto --go_out=src/
protoc protocol/rpc.proto --go_out=plugins=grpc:src/
