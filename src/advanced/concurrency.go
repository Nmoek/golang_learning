/*
 * @file: concurrency.go
 * @brief: 并发预定义的函数
 * @author: Kewin Li
 * @date: 2023-04-04
 */

package concurrency_test

type WebsiteChecker func(string) bool

// @func: CheckWebsites
// @brief: 检查每一个传入的url合法性
// @author: Kewin Li
// @param: WebsiteChecker wc
// @param: []string urls
// @return map
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
