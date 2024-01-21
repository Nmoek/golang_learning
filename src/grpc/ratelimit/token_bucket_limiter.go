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
// @Description: 限流算法-令牌桶
type TokenBucketLimiter struct {
	interval  time.Duration // 多久生产一个令牌
	buckets   chan struct{}
	closeChan chan struct{}
	closeOne  sync.Once
}

func NewTokenBucketLimiter(interval time.Duration) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		interval:  interval,
		buckets:   make(chan struct{}, 20),
		closeChan: make(chan struct{}),
	}
}

func (t *TokenBucketLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {

	ticker := time.NewTicker(t.interval)
	go func() {
		for {
			select {
			// 计时到期
			case <-ticker.C:

				select {
				// 发放令牌
				case t.buckets <- struct{}{}:

				// 令牌桶已经满
				default:
				}

			// 关闭令牌发放
			case <-t.closeChan:
				return
			}
		}
	}()

	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		select {
		// 拿到令牌
		case <-t.buckets:
			return handler(ctx, req)
		// 没拿到令牌阻塞等待直到超时
		//case <-ctx.Done():
		//	return nil, ctx.Err()

		// 没拿到令牌就返回
		default:
			return nil, status.Errorf(codes.ResourceExhausted, "触发限流")

		}
	}
}

func (t *TokenBucketLimiter) Close() error {
	// 只能关闭一次chan 否则触发panic
	t.closeOne.Do(func() {
		close(t.closeChan)
	})
	return nil
}
