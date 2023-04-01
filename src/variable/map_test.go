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

type Dictionary map[string]string

// @func: Search
// @brief: 哈希表抽象为方法返回查询结果
// @author: Kewin Li
// @param: map[string]string dictionary
// @param: string key
// @return string
func (dictionary Dictionary) Search(key string) string {
	return dictionary[key]
}

// @func: Search
// @brief: 返回哈希查询结果
// @author: Kewin Li
// @param: map[string]string dictionary
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

	// 通过函数调用查询结果
	t.Run("Search map result", func(t *testing.T) {

		got := Search(dictionary, "test")
		want := "this my test"
		errorMsg(t, got, want)
	})

	// 通过方法查询结果
	t.Run("Search function test", func(t *testing.T) {

		dic := Dictionary{"test2": "this my test"}
		got := dic.Search("test2")
		want := "this my test"

		errorMsg(t, got, want)

	})

}

func errorMsg(t *testing.T, got string, want string) {

	t.Helper()
	if got != want {
		t.Errorf("got= %s want= %s \n", got, want)
	}

}
