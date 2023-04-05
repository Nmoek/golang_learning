/*
 * @file: select_test.go
 * @brief: select测试学习
 * @author: Kewin Li
 * @date: 2023-04-05
 */

package select_test

import (
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

// @func: TestRacer
// @brief: 测试某个URL返回更快
// @author: Kewin Li
// @param: *testing.T t
func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	quickServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	quickUrl := quickServer.URL

	// 同步阻塞发出http请求
	t.Run("sync send http request!", func(t *testing.T) {
		got := Racer(slowUrl, quickUrl)
		want := quickUrl
		if got != want {
			t.Errorf("got=%s  want=%s \n", got, want)
		}
	})

	// 并发发出http请求
	t.Run("concurrency send http request!", func(t *testing.T) {
		got := RacerWithSelect(slowUrl, quickUrl)
		want := quickUrl
		if got != want {
			t.Errorf("got=%s  want=%s \n", got, want)
		}
	})

	slowServer.Close()
	quickServer.Close()
}
