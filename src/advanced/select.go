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
	"net/http/httptest"
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

// @func: Racer
// @brief: 返回响应较快的URL
// @author: Kewin Li
// @param: string a
// @param: string b
// @return string
func Racer(a string, b string) string {
	durationA := GetDuration(a)
	durationB := GetDuration(b)

	if durationA < durationB {
		return a
	}

	return b
}

// @func: CreateHTTPServer
// @brief: 创建一个HTTP测试服务器对象
// @author: Kewin Li
// @param: time.Duration delay
// @return *httptest.Server
func CreateHTTPServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}

func main() {

	slowServer := CreateHTTPServer(20 * time.Millisecond)
	quickServer := CreateHTTPServer(0 * time.Microsecond)

	defer slowServer.Close()
	defer quickServer.Close()

	slowUrl := slowServer.URL
	quickUrl := quickServer.URL

	fmt.Printf("slowUrl=%s  quickUrl=%s \n", slowUrl, quickUrl)

	result := Racer(slowUrl, quickUrl)

	fmt.Printf("result=%s \n", result)

}
