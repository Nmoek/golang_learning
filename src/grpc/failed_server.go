package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FailedServer struct {
	UnimplementedUserServiceServer
	Name string
}

func (s *FailedServer) GetById(ctx context.Context, request *GetByIdRequest) (*GetByIdResponse, error) {
	fmt.Printf("收到来自客户端的调用, %v\n", request.Id)

	return nil, status.Errorf(codes.Unavailable, "已被熔断")
}
