// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todolist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodolistServiceClient is the client API for TodolistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodolistServiceClient interface {
	CreateTodolist(ctx context.Context, in *CreateTodolistRequest, opts ...grpc.CallOption) (*CreateTodolistResponse, error)
	ListTodoList(ctx context.Context, in *ListTodolistRequest, opts ...grpc.CallOption) (*ListTodolistResponse, error)
}

type todolistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodolistServiceClient(cc grpc.ClientConnInterface) TodolistServiceClient {
	return &todolistServiceClient{cc}
}

func (c *todolistServiceClient) CreateTodolist(ctx context.Context, in *CreateTodolistRequest, opts ...grpc.CallOption) (*CreateTodolistResponse, error) {
	out := new(CreateTodolistResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodolistService/CreateTodolist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todolistServiceClient) ListTodoList(ctx context.Context, in *ListTodolistRequest, opts ...grpc.CallOption) (*ListTodolistResponse, error) {
	out := new(ListTodolistResponse)
	err := c.cc.Invoke(ctx, "/todolist.TodolistService/ListTodoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodolistServiceServer is the server API for TodolistService service.
// All implementations must embed UnimplementedTodolistServiceServer
// for forward compatibility
type TodolistServiceServer interface {
	CreateTodolist(context.Context, *CreateTodolistRequest) (*CreateTodolistResponse, error)
	ListTodoList(context.Context, *ListTodolistRequest) (*ListTodolistResponse, error)
	mustEmbedUnimplementedTodolistServiceServer()
}

// UnimplementedTodolistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodolistServiceServer struct {
}

func (UnimplementedTodolistServiceServer) CreateTodolist(context.Context, *CreateTodolistRequest) (*CreateTodolistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodolist not implemented")
}
func (UnimplementedTodolistServiceServer) ListTodoList(context.Context, *ListTodolistRequest) (*ListTodolistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodoList not implemented")
}
func (UnimplementedTodolistServiceServer) mustEmbedUnimplementedTodolistServiceServer() {}

// UnsafeTodolistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodolistServiceServer will
// result in compilation errors.
type UnsafeTodolistServiceServer interface {
	mustEmbedUnimplementedTodolistServiceServer()
}

func RegisterTodolistServiceServer(s grpc.ServiceRegistrar, srv TodolistServiceServer) {
	s.RegisterService(&TodolistService_ServiceDesc, srv)
}

func _TodolistService_CreateTodolist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodolistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServiceServer).CreateTodolist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodolistService/CreateTodolist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServiceServer).CreateTodolist(ctx, req.(*CreateTodolistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodolistService_ListTodoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodolistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodolistServiceServer).ListTodoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todolist.TodolistService/ListTodoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodolistServiceServer).ListTodoList(ctx, req.(*ListTodolistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodolistService_ServiceDesc is the grpc.ServiceDesc for TodolistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodolistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todolist.TodolistService",
	HandlerType: (*TodolistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodolist",
			Handler:    _TodolistService_CreateTodolist_Handler,
		},
		{
			MethodName: "ListTodoList",
			Handler:    _TodolistService_ListTodoList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todolist/todolist.proto",
}
