package cronjob

import (
	"context"
	"testing"
	"time"
)

// @func: TestTicker
// @date: 2023-12-23 22:20:24
// @brief: 使用time.Ticker控制定时任务
// @author: Kewin Li
// @param t
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer ticker.Stop()
	for {

		select {
		// 每隔1s会触发一个信号，能够通过channel取出
		case now := <-ticker.C:
			t.Log(now.UnixMilli(), "ms")
		case <-ctx.Done():
			t.Log("退出")
			// select中使用break无法退出循环
			goto EXIT
		}
	}
EXIT:
}
