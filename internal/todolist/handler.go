package todolist

import (
	"context"
	"fmt"
	"time"

	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
	"github.com/rs/xid"
)

var s = make([]*tdlpb.TodolistResponseData, 0)

type handler struct {
	tdlpb.UnimplementedTodolistServiceServer
}

func NewHandler() *handler {
	fmt.Println("-------------call handler---------------")
	return &handler{}
}

func (h handler) CreateTodolist(ctx context.Context, in *tdlpb.CreateTodolistRequest) (*tdlpb.CreateTodolistResponse, error) {
	resp := tdlpb.CreateTodolistResponse{
		Data: &tdlpb.TodolistResponseData{
			Title:       in.Title,
			Id:          xid.NewWithTime(time.Now()).String(),
			Description: in.Description,
			Status:      in.Status,
			CreatedDate: in.CreatedDate,
		},
	}
	s = append(s, resp.Data)

	return &resp, nil
}

func (h handler) ListTodolist(ctx context.Context, in *tdlpb.ListTodolistRequest) (*tdlpb.ListTodolistResponse, error) {
	return &tdlpb.ListTodolistResponse{Data: s}, nil
}
