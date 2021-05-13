#!/bin/bash


protoc -I ./api/proto \
--go_out ./api/proto --go_opt paths=source_relative \
--go-grpc_out ./api/proto --go-grpc_opt paths=source_relative \
--grpc-gateway_out ./api/proto --grpc-gateway_opt paths=source_relative \
./api/proto/todolist/todolist.proto