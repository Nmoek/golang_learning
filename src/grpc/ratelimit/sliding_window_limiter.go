package ratelimit

import (
	"context"
	"github.com/liyue201/gostl/ds/priorityqueue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	window    time.Duration
	queue     *priorityqueue.PriorityQueue[time.Time] // 小顶堆
	threshold int
	lock      sync.Mutex
}

func NewSlidingWindowLimiter(window time.Duration, threshold int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		window:    window,
		threshold: threshold,
		queue: priorityqueue.New[time.Time](func(a, b time.Time) int {
			if a.Before(b) {
				return -1
			} else if a.After(b) {
				return 1
			}
			return 0
		}),
	}
}

func (s *SlidingWindowLimiter) BuildServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		s.lock.Lock()

		// 1. 直接判断当窗口是否触发限流
		if s.queue.Size() < s.threshold {
			s.lock.Unlock()
			return handler(ctx, req)
		}

		// 2. 移动窗口
		windowStart := time.Now().Add(-s.window)
		for {
			minReq := s.queue.Top()
			if minReq.Before(windowStart) {
				s.queue.Pop()
			} else {
				break
			}
		}

		if s.queue.Size() < s.threshold {
			s.queue.Push(time.Now())
			s.lock.Unlock()
			// 3. 执行请求
			return handler(ctx, req)
		}

		s.lock.Unlock()
		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
	}
}
