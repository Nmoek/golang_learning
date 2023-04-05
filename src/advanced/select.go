/*
 * @file: select.go
 * @brief: select自测示例
 * @author: Kewin Li
 * @date: 2023-04-05
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

type _result struct {
	url      string
	duration float64
}

// @func: GetDuration
// @brief: 获取http请求返回时间间隔
// @author: Kewin Li
// @param: string url
// @return float64
func GetDuration(url string) float64 {

	start := time.Now()
	http.Get(url)

	return float64(time.Since(start))

}

func main() {

	a := "http://github.com"
	b := "http://baidu.com"

	urls := []string{a, b}
	channelResult := make(chan _result)

	for _, url := range urls {

		go func(u string) {
			channelResult <- _result{u, GetDuration(u)}
		}(url)

	}

	for i := 0; i < len(urls); i++ {

		go func() {
			r := <-channelResult //疑似阻塞等待

			fmt.Printf("url=%s duration=%.2f ms \n", r.url, r.duration/1000000.0)
		}()
	}

	time.Sleep(11 * time.Second)
}
