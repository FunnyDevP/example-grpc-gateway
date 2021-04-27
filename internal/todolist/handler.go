package todolist

import (
	"context"

	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
	"google.golang.org/grpc"
)

type handler struct {
	tdlpb.UnimplementedTodolistServiceServer
}

func NewHandler() *handler {
	return &handler{}
}

func (h handler) CreateTodolist(ctx context.Context, in *tdlpb.CreateTodolistRequest, opts ...grpc.CallOption) (*tdlpb.CreateTodolistResponse, error)  {

}




