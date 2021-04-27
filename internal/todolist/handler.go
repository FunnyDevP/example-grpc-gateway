package todolist

import (
	"context"
	"time"

	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
)

type handler struct {
	tdlpb.UnimplementedTodolistServiceServer
}

func NewHandler() *handler {
	return &handler{}
}

func (h handler) CreateTodolist(ctx context.Context, in *tdlpb.CreateTodolistRequest) (*tdlpb.CreateTodolistResponse, error)  {
	return &tdlpb.CreateTodolistResponse{
		Data: &tdlpb.TodolistResponseData{
			Id:          "1234",
			Description: "Hello description",
			Status:      "pending",
			CreatedDate: time.Now().String(),
		},
	},nil
}




