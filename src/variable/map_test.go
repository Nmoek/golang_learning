/*
 * @file:maps_test.go
 * @brief: maps的测试学习
 * @author: Kewin Li
 * @date:2023-04-01
 */

package main

import (
	"errors"
	"testing"
)

type Dictionary map[string]string

var errorNotFound = errors.New("can not find result!")

// @func: Search
// @brief: 哈希表抽象为方法返回查询结果
// @author: Kewin Li
// @param: map[string]string dictionary
// @param: string key
// @return string
func (dictionary Dictionary) Search(key string) (string, error) {

	result, ok := dictionary[key]

	if !ok {
		return "", errorNotFound
	}

	return result, nil
}

// @func: Search
// @brief: 返回哈希查询结果
// @author: Kewin Li
// @param: map[string]string dictionary
// @return string
func Search(dictionary map[string]string, key string) (string, error) {

	result, ok := dictionary[key]
	if !ok {
		return "", errorNotFound
	}

	return result, nil
}

// @func: TestSearch
// @brief: 测试键值搜索
// @author: Kewin Li
// @param: *testing.T t
func TestSearch(t *testing.T) {

	dictionary := map[string]string{"test": "this my test"}

	// 通过函数调用查询结果
	t.Run("Search map result", func(t *testing.T) {

		got, _ := Search(dictionary, "test")
		want := "this my test"
		checkResult(t, got, want)
	})

	// 通过方法查询结果
	t.Run("Search function test", func(t *testing.T) {

		dic := Dictionary{"test2": "this my test"}
		got, _ := dic.Search("test2")
		want := "this my test"

		checkResult(t, got, want)

	})

	//重点: 非法key的处理
	t.Run("unknow key", func(t *testing.T) {

		dic := Dictionary{"test3": "this my test"}
		_, err := dic.Search("unknow")

		errorMsg(t, err, errorNotFound)
	})

}

// @func: checkResult
// @brief: 检查查询结果
// @author: Kewin Li
// @param: *testing.T t
// @param: string got
// @param: string want
func checkResult(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got=%s, want=%s", got, want)
	}
}

// @func: errorMsg
// @brief: 出错断言
// @author: Kewin Li
// @param: *testing.T t
// @param: error got
// @param: error want
func errorMsg(t *testing.T, got error, want error) {

	t.Helper()
	if got != want {
		t.Errorf("got err=%s, want err=%s \n", got, want)
	}

}
