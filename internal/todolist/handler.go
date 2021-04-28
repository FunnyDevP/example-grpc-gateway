package todolist

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	tdlpb "github.com/FunnyDevP/example-grpc-gateway/api/proto/todolist"
	"github.com/rs/xid"
	"google.golang.org/grpc"
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

	setHttpStatusCode(ctx, http.StatusCreated)
	return &resp, nil
}

func (h handler) ListTodolist(ctx context.Context, in *tdlpb.ListTodolistRequest) (*tdlpb.ListTodolistResponse, error) {
	setHttpStatusCode(ctx, http.StatusOK)
	return &tdlpb.ListTodolistResponse{Data: s}, nil
}

func (h handler) EditTodolist(ctx context.Context, in *tdlpb.EditTodolistRequest) (*tdlpb.EditTodolistResponse, error) {
	data := new(tdlpb.TodolistResponseData)
	for _, rs := range s {
		if in.TodolistId == rs.Id {
			rs.Title = in.TodolistBody.Title
			rs.Description = in.TodolistBody.Description
			rs.Status = in.TodolistBody.Status
			rs.CreatedDate = in.TodolistBody.CreatedDate
			data = rs
			break
		}
	}
	setHttpStatusCode(ctx, http.StatusOK)
	return &tdlpb.EditTodolistResponse{Data: data}, nil
}

func setHttpStatusCode(ctx context.Context, sc int) {
	_ = grpc.SetHeader(ctx, map[string][]string{
		"x-http-code": {strconv.Itoa(sc)},
	})
}
