/*
 * @file:maps_test.go
 * @brief: maps的测试学习
 * @author: Kewin Li
 * @date:2023-04-01
 */

package main

import (
	"testing"
)

// @func: Search
// @brief: 返回哈希映射的结果
// @author: Kewin Li
// @param: map[string]string dictionary
// @param: string key
// @return string
func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

// @func: TestSearch
// @brief: 测试键值搜索
// @author: Kewin Li
// @param: *testing.T t
func TestSearch(t *testing.T) {

	dictionary := map[string]string{"test": "this my test"}

	got := Search(dictionary, "test")
	want := "this my test"

	if got != want {
		t.Errorf("got= %s want= %s", got, want)
	}

}
