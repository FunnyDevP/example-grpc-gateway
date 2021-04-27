package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
	"github.com/FunnyDevP/example-grpc-gateway/internal/todolist"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln("Failed to listen:",err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the todolist service to the server
	tdlpb.RegisterTodolistServiceServer(s,todolist.NewHandler())
	// Serve gRPC server
	log.Println("Serving gRPC on localhost:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:",err)
	}

	gwmux := runtime.NewServeMux()
	// register http handler
	if err := tdlpb.RegisterTodolistServiceHandler(context.Background(),gwmux,conn); err != nil {
		log.Fatalln("Failed to register gateway:",err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
