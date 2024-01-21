package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_test/ratelimit"
)

// LimiterUserServer
// @Description: 业务级别限流-装饰器
type LimiterUserServer struct {
	limiter ratelimit.Limiter
	UserServiceServer
	Name string
}

func NewLimiterUserServer(limiter ratelimit.Limiter) *LimiterUserServer {
	return &LimiterUserServer{
		limiter: limiter,
		Name:    "biz_limiter_test",
	}
}

func (l *LimiterUserServer) GetById(ctx context.Context, req *GetByIdRequest) (*GetByIdResponse, error) {

	key := fmt.Sprintf("limiter:user:get_by_id:%d", req.Id)
	limited, err := l.limiter.Limit(ctx, key)
	if err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
	}

	if limited {
		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")

	}

	return l.UserServiceServer.GetById(ctx, req)
}
