/*
 * @file:maps_test.go
 * @brief: maps的"增删改查"测试学习
 * @author: Kewin Li
 * @date:2023-04-01
 */

package main

import (
	"fmt"
	"testing"
)

type Dictionary map[string]string

type DictionaryError string

const (
	errorNotFound    = DictionaryError("can not find result!")
	errorKeyNull     = DictionaryError("add key is null!")
	errorKeyExist    = DictionaryError("add key is existed!")
	errorKeyNotExist = DictionaryError("key is not existed!")
)

// @func: Error
// @brief: 重载错误打印接口
// @author: Kewin Li
// @receiver: DictionaryError e
// @return string
func (e DictionaryError) Error() string {
	return string(e)
}

// @func: Search
// @brief: 方法返回查询结果
// @author: Kewin Li
// @param: map[string]string dictionary
// @param: string key
// @return string
func (dictionary Dictionary) Search(key string) (string, error) {

	result, ok := dictionary[key]

	if ok != true {
		return "", errorNotFound
	}

	return result, nil
}

// @func: Search
// @brief: h函数返回查询结果
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
		got, err := dic.Search("unknow")

		//打印一下，传入非法key后会得到什么
		fmt.Printf("map result info: ")

		fmt.Printf("got type=%T val=%v\n", got, got)

		errorMsg(t, err, errorNotFound)
	})

}

// @func: AddKey
// @brief: 函数添加键值
// @author: Kewin Li
// @param: Dictionary d
// @param: string key
// @param: string result
// @return string
// @return error
func AddKey(d Dictionary, key string, result string) error {
	if key == "" {
		return errorKeyNull
	}

	_, err := d.Search(key)

	switch err {
	case errorNotFound:
		{
			d[key] = result
		}
	case nil:
		{
			return errorKeyExist
		}

	default:
		{
			return err
		}
	}

	return nil
}

// @func: AddKey
// @brief: 方法添加键值
// @author: Kewin Li
// @receiver: Dictionary d
// @param: string key
// @param: string result
// @return error
func (d Dictionary) AddKey(key string, result string) error {
	if key == "" {
		return errorKeyNull
	}

	val, _ := d.Search(key)
	if val != "" {
		return errorKeyExist
	}

	d[key] = result

	return nil
}

// @func: TestAdd
// @brief: 测试键值添加
// @author: Kewin Li
// @param: *testing.T t
func TestAdd(t *testing.T) {

	//通过函数调用添加键值
	t.Run("test add 1", func(t *testing.T) {
		dic := Dictionary{}

		err := AddKey(dic, "test add", "this a add")
		if err != nil {
			t.Fatalf("err= %s\n", err.Error())
		}

		got, _ := dic.Search("test add")
		want := "this a add"

		checkResult(t, got, want)
	})

	// 通过方法调用添加键值
	t.Run("test add 2", func(t *testing.T) {

		dic := Dictionary{}
		err := dic.AddKey("test add2", "this a add2")
		if err != nil {
			t.Fatalf("err=%s \n", err.Error())
		}

		got, _ := dic.Search("test add2")
		want := "this a add2"

		checkResult(t, got, want)

	})

	//尝试添加一个相同的key
	t.Run("test add exist key", func(t *testing.T) {

		key := "test3"
		val := "this my test3"
		dic := Dictionary{key: val}

		err := dic.AddKey("test3", "oh no")

		if err != nil {
			t.Fatalf("key=%s err=%s \n", key, err.Error())
		}

	})

}

// @func: Update
// @brief: 函数修改键值对
// @author: Kewin Li
// @param: Dictionary d
// @param: string key
// @param: string val
// @return error
func Update(d Dictionary, key string, val string) error {

	_, err := d.Search(key)

	if err != nil {
		return errorKeyNotExist
	}

	d[key] = val

	return nil
}

// @func: Update
// @brief: 方法修改键值对
// @author: Kewin Li
// @receiver: Dictionary d
// @param: string key
// @param: string val
// @return error
func (d Dictionary) Update(key string, val string) error {
	_, err := d.Search(key)

	// key不存在不能进行修改动作
	if err != nil {
		return errorKeyNotExist
	}

	d[key] = val

	return nil
}

// @func: TestUpdate
// @brief: 测试键值修改
// @author: Kewin Li
// @param: *testing.T t
func TestUpdate(t *testing.T) {

	key := "test"
	val := "this my test"
	dic := Dictionary{key: val}

	//通过函数调用修改键值对
	t.Run("test update 1", func(t *testing.T) {

		val = "this my update1"

		err := Update(dic, key, val)
		if err != nil {
			t.Errorf("key='%s' val='%s' err='%s' \n", key, val, err.Error())
		}

		result, _ := dic.Search(key)

		checkResult(t, result, val)

	})

	// 通过方法调用修改键值对
	t.Run("test update 2", func(t *testing.T) {

		val = "this my update2"

		err := dic.Update("test1", val)
		if err != nil {
			t.Errorf("key='%s' val='%s' err='%s' \n", key, val, err.Error())
		}

		result, _ := dic.Search(key)

		checkResult(t, result, val)

	})

}

// @func: DeleteKey
// @brief: 函数删除键值对
// @author: Kewin Li
// @param: Dictionary d
// @param: string key
// @return error
func DeleteKey(d Dictionary, key string) error {

	_, err := d.Search(key)

	if err != nil {
		return err
	}

	delete(d, key)

	return nil

}

// @func: DeleteKey
// @brief: 方法删除键值对
// @author: Kewin Li
// @receiver: Dictionary d
// @param: string key
// @return error
func (d Dictionary) DeleteKey(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return err
	}

	delete(d, key)

	return nil

}

// @func: TestDelete
// @brief: 测试键值对删除
// @author: Kewin Li
// @param: *testing.T t
func TestDelete(t *testing.T) {

	key := "test"
	val := "this a test"
	dic := Dictionary{}

	for i := 0; i < 10; i++ {
		new_key := fmt.Sprintf("%s%d", key, i)
		new_val := fmt.Sprintf("%s%d", val, i)

		err := dic.AddKey(new_key, new_val)
		if err != nil {
			t.Fatalf("key='%s' val='%s' err='%s' \n", new_key, new_val, err.Error())
		}
	}

	fmt.Printf("\ndic=%v \n", dic)

	//函数调用删除键值对
	t.Run("test delete 1", func(t *testing.T) {

		del_key := "test0"

		err := DeleteKey(dic, del_key)
		if err != nil {
			t.Fatalf("del_key='%s' err='%s' \n", del_key, err.Error())
		}

		fmt.Printf("\ndic=%v \n", dic)

	})

	//方法调用删除键值对
	t.Run("test delete 2", func(t *testing.T) {
		del_key := "test1"

		err := dic.DeleteKey(del_key)
		if err != nil {
			t.Fatalf("del_key='%s' err='%s' \n", del_key, err.Error())
		}

		fmt.Printf("\ndic=%v \n", dic)

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
