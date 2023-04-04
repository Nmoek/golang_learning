/*
 * @file: concurrency_test.go
 * @brief: Go的并发测试学习
 * @author: Kewin Li
 * @date: 2023-04-04
 */
package concurrency_test

import (
	"reflect"
	"testing"
)

// @func: mockWebsiteCHecker
// @brief:
// @author: Kewin Li
// @param: string url
// @return bool
func mockWebsiteChecker(url string) bool {
	if url == "http://baidu.com" {
		return true
	}

	return false
}

// @func: TestWebsiteChecker
// @brief: 测试检查url的状态
// @author: Kewin Li
// @param: *testing.T t
func TestWebsiteChecker(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://baidu.com",
	}

	results := CheckWebsites(mockWebsiteChecker, websites)

	got := len(websites)
	want := len(results)
	if got != want {
		t.Fatalf("got=%v want=%v \n", got, want)
	}

	expectResults := map[string]bool{
		"http://google.com": false,
		"http://baidu.com":  true,
	}

	if !reflect.DeepEqual(results, expectResults) {
		t.Fatalf("res=%v  expectRes=%v \n", results, expectResults)
	}

}
