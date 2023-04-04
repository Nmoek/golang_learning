/*
 * @file: concurrency_test.go
 * @brief: Go的并发测试学习
 * @author: Kewin Li
 * @date: 2023-04-04
 */
package concurrency_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type WebsiteChecker func(string) bool

type _result struct {
	string
	bool
}

// @func: CheckWebsites
// @brief: 检查urls列表的合法性
// @author: Kewin Li
// @param: WebsiteChecker wc
// @param: []string urls
// @return map
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan _result)

	//向各个子线程发送数据
	for _, url := range urls {
		//匿名函数直接开启一个新的goroutine
		go func(u string) {
			fmt.Printf("before r.string=%s \n", u)
			resultChannel <- _result{u, wc(u)}
		}(url)
	}

	//从各个子线程回收数据
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		fmt.Printf("after r.string=%s r.bool=%v \n", r.string, r.bool)
		results[r.string] = r.bool
	}

	return results
}

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

// @func: slowWebsiteChecker
// @brief: 模拟查询url的耗时
// @author: Kewin Li
// @param: string _
// @return bool
func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// @func: BenchmarkCheckWebsites
// @brief: 基准测试url合法查询, 性能粗侧
// @author: Kewin Li
// @param: *testing.B b
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
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
		"http://baidu.co":   true,
	}

	if !reflect.DeepEqual(results, expectResults) {
		t.Fatalf("res=%v  expectRes=%v \n", results, expectResults)
	}

}
