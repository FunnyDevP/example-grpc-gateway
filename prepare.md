1. define service and message on protobuf file

```
rpc ListTodoList(ListTodolistRequest) returns(ListTodolistResponse) {
    option (google.api.http) = {
        get: "/api/v1/todolists"
    };
}
message ListTodolistRequest {
}

message ListTodolistResponse {
    repeated TodolistResponseData data = 1;
}
```

2. generate _.pb.go and _.pb.gw.go
```
./gen_proto.sh
```

3. implementation service
```
func (h handler) ListTodoList(ctx context.Context, in *tdlpb.ListTodolistRequest) (*tdlpb.ListTodolistResponse, error)  {
	setHttpStatusCode(ctx, http.StatusOK)
	return &tdlpb.ListTodolistResponse{Data: s}, nil
}
```

4. customization error response 
```

md, _ := metadata.FromIncomingContext(ctx)

if len(md["authorization-time"]) == 0 {
    return nil, status.New(codes.InvalidArgument,"authorization-time must be specify").Err()
}
	
type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func customHttpResponseErr(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	s, _ := status.FromError(err)

	httpCode := runtime.HTTPStatusFromCode(s.Code())

	errResp := ErrorResponse{Error: Error{
		Code: httpCode,
		Message: s.Message(),
	}}


	b, _ := marshaler.Marshal(errResp)
	writer.WriteHeader(httpCode)
	_, _ = writer.Write(b)
}
```

