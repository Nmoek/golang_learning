// Package channel
// @Description: channel测试练习
package channel

import (
	"fmt"
	"testing"
	"time"
)

// @func: TestChannel
// @date: 2023-12-17 01:26:38
// @brief: channal的基本使用
// @author: Kewin Li
// @param t
func TestChannel(t *testing.T) {

	ch := make(chan int)
	i := 0

	go func() {
		for {
			time.Sleep(1 * time.Second)

			data, ok := <-ch
			if ok {
				fmt.Printf("[OUT]============ %v ms, data:%d \n", time.Now().UnixMilli(), data)

			}
		}

	}()

	for {

		// 非缓存队列时会将goroutine阻塞
		ch <- i
		fmt.Printf("[IN]============ %v ms, data:%d \n", time.Now().UnixMilli(), i)

		i++
		time.Sleep(time.Millisecond * 500)
	}

}
