package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"

	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
	td "github.com/FunnyDevP/example-grpc-gateway/internal/todolist"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
	tdlpb.RegisterTodolistServiceServer(s,td.NewHandler())
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

	gwmux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(httpResponseCodeModifier),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaller runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {

		}),
	)

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

func httpResponseCodeModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	key := "x-http-code"
	// set http status code
	if vals := md.HeaderMD.Get(key); len(vals) > 0 {
		code , err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}

		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD,key)
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}
	return nil
}
