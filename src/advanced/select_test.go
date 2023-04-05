/*
 * @file: select_test.go
 * @brief: select测试学习
 * @author: Kewin Li
 * @date: 2023-04-05
 */

package select_test

import (
	"net/http"
	"testing"
	"time"
)

// @func: Racer
// @brief: 某个URL返回更快
// @author: Kewin Li
// @param: string slowUrl
// @param: string quickUrl
func Racer(a string, b string) string {

	start1 := time.Now()
	http.Get(a)
	duration1 := time.Since(start1)

	start2 := time.Now()
	http.Get(b)
	duration2 := time.Since(start2)

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

	slowUrl := "http://github.com"
	quickUrl := "http://baidu.com"

	got := Racer(slowUrl, quickUrl)
	want := quickUrl

	if got != want {
		t.Errorf("got=%s want=%s \n", got, want)
	}

}