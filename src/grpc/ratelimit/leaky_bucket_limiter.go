package ratelimit

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

// LeakyBucketLimiter
// @Description: 限流算法-漏桶
type LeakyBucketLimiter struct {
	interval  time.Duration // 多久生产一个令牌
	closeChan chan struct{}
	closeOne  sync.Once
}

func NewLeakyBucketLimiter(interval time.Duration) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		interval:  interval,
		closeChan: make(chan struct{}),
	}
}

func (l *LeakyBucketLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	ticker := time.NewTicker(l.interval)

	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		select {
		// 拿到令牌
		case <-ticker.C:
			return handler(ctx, req)
		// 关闭令牌发放
		case <-l.closeChan:
			return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
		// 没拿到令牌阻塞等待直到超时
		//case <-ctx.Done():
		//	return nil, ctx.Err()

		// 没拿到令牌就返回
		default:
			return nil, status.Errorf(codes.ResourceExhausted, "触发限流")

		}
	}
}

func (l *LeakyBucketLimiter) Close() error {
	// 只能关闭一次chan 否则触发panic
	l.closeOne.Do(func() {
		close(l.closeChan)
	})
	return nil
}
