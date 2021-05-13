1. define service and message on .proto file
2. clone repo https://github.com/googleapis/googleapis and copy some dependencies into our proto file structure.
2. generate *.pb, *_grpc.pb.go and *.pb.gw.go

**using protoc**
```
   protoc -I ./api/proto \ 
   
   --go_out ./api/proto --go_opt paths=source_relative \
   
   --go-grpc_out ./api/proto --go-grpc_opt paths=source_relative \
   
   --grpc-gateway_out ./api/proto --grpc-gateway_opt paths=source_relative \
   
   ./api/proto/todolist/todolist.proto
```
**using buf**
```
 buf generate --template scripts/proto/todolist.gen.yaml --path api/proto/todolist/todolist.proto

```

