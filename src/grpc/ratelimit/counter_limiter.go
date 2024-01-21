package ratelimit

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync/atomic"
)

// CounterLimiter
// @Description: 限流算法-计数器算法实现
type CounterLimiter struct {
	cnt       atomic.Int32
	threshold int32
}

func (c *CounterLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		cnt := c.cnt.Add(1)
		defer func() {
			c.cnt.Add(-1)
		}()

		// 小于阈值就执行请求
		if cnt <= c.threshold {
			return handler(ctx, req)
		}

		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
	}
}
