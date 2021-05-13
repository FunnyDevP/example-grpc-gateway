package todolist

import (
	"context"
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

func setHttpStatusCode(ctx context.Context, sc int) {
	_ = grpc.SetHeader(ctx, map[string][]string{
		"x-http-code": {strconv.Itoa(sc)},
	})
}
