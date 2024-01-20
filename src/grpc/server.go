package grpc

import (
	"context"
	"fmt"
)

type Server struct {
	UnimplementedUserServiceServer
	Name string
}

func (s *Server) GetById(ctx context.Context, request *GetByIdRequest) (*GetByIdResponse, error) {
	fmt.Printf("收到来自客户端的调用, %v\n", request.Id)

	return &GetByIdResponse{
		User: &User{
			Id:    666,
			Name:  "ljk from" + s.Name,
			Addrs: "hangzhou",
		},
	}, nil
}
