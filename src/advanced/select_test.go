/*
 * @file: select_test.go
 * @brief: select测试学习
 * @author: Kewin Li
 * @date: 2023-04-05
 */

package select_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// @func: GetDuration
// @brief: 获取请求URL间隔
// @author: Kewin Li
// @param: string url
// @return flaot64 ms
func GetDuration(url string) float64 {

	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return float64(duration) / 1000000.0
}

// @func: ping
// @brief: 辅助函数, 发起http请求
// @author: Kewin Li
// @param: string url
// @return chan
func ping(url string) chan bool {

	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}

// @func: RacerWithSelect
// @brief: 同时开启多个goroutine执行并发with超时机制
// @author: Kewin Li
// @param: string a
// @param: string b
// @return string
func RacerWithTimeout(a string, b string, timeout time.Duration) (string, error) {

	select {
	case <-ping(a):
		return a, nil

	case <-ping(b):
		return b, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("time out for %s and %s", a, b)
	}

	return "", nil
}

// @func: RacerWithSelect
// @brief: 同时开启多个goroutine执行并发
// @author: Kewin Li
// @param: string a
// @param: string b
// @return string
func RacerWithSelect(a string, b string) string {

	select {
	case <-ping(a):
		return a

	case <-ping(b):
		return b

	}

	return ""
}

// @func: Racer
// @brief: 某个URL返回更快
// @author: Kewin Li
// @param: string slowUrl
// @param: string quickUrl
func Racer(a string, b string) string {

	duration1 := GetDuration(a)
	duration2 := GetDuration(b)

	if duration1 < duration2 {
		return a
	}

	return b
}

// @func: CreateHTTPServer
// @brief: 创建http测试服务器
// @author: Kewin Li
// @param: time.Duration d
// @return *http.Server
func CreateTestHTTPServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

// @func: TestRacer
// @brief: 测试某个URL返回更快
// @author: Kewin Li
// @param: *testing.T t
func TestRacer(t *testing.T) {
	// 同步阻塞发出http请求
	t.Run("sync send http request!", func(t *testing.T) {
		server1 := CreateTestHTTPServer(20 * time.Microsecond)
		server2 := CreateTestHTTPServer(0 * time.Microsecond)
		defer server1.Close()
		defer server2.Close()

		got := Racer(server1.URL, server2.URL)
		want := server2.URL
		if got != want {
			t.Errorf("got=%s  want=%s \n", got, want)
		}

	})

	// 并发发出http请求
	t.Run("concurrency send http request!", func(t *testing.T) {
		server1 := CreateTestHTTPServer(20 * time.Microsecond)
		server2 := CreateTestHTTPServer(0 * time.Microsecond)
		defer server1.Close()
		defer server2.Close()

		got := RacerWithSelect(server1.URL, server2.URL)
		want := server2.URL
		if got != want {
			t.Errorf("got=%s  want=%s \n", got, want)
		}
	})

	// 测试select超时机制
	t.Run("select with timeout send http request!", func(t *testing.T) {

		server1 := CreateTestHTTPServer(12 * time.Second)
		server2 := CreateTestHTTPServer(11 * time.Second)
		defer server1.Close()
		defer server2.Close()

		winner, err := RacerWithTimeout(server1.URL, server2.URL, 5*time.Second)
		want := server2.URL

		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}

		if winner != want {
			t.Errorf("got=%s want=%s \n", winner, want)
		}

	})

}
