package ratelimit

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

// CounterLimiter
// @Description: 限流算法-固定窗口算法实现
type FixedWindowLimiter struct {
	window     time.Duration //每个窗口的大小
	lastWindow time.Time     //上一个窗口的起始点
	cnt        int           // 当前窗口的请求数量
	threshold  int           // 限流阈值
	lock       sync.Mutex
}

func NewFixedWindowLimiter(window time.Duration, threshold int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		cnt:        0,
		window:     window,
		threshold:  threshold,
		lastWindow: time.Now(),
	}
}

func (f *FixedWindowLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		f.lock.Lock()

		// 1. 判断当前窗口是否已经过期
		// 		1.1 过期新开窗口
		now := time.Now()
		if now.After(f.lastWindow.Add(f.window)) {
			f.lastWindow = now
			f.cnt = 0
		}
		f.cnt += 1
		cnt := f.cnt
		f.lock.Unlock()

		if cnt <= f.threshold {
			return handler(ctx, req)
		}

		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
	}
}
