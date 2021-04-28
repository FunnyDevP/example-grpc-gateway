module github.com/FunnyDevP/example-grpc-gateway

// +build tools

go 1.15

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/rs/xid v1.3.0
	google.golang.org/genproto v0.0.0-20210224155714-063164c882e6
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0
)
