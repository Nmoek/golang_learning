// Package main
// @Description: 测试各种local cache用法
package main

import (
	"fmt"
	"github.com/coocood/freecache"
	"os"
	"runtime/debug"
)

// 100 M
const totalCacheSize = 100 * 1024 * 1024

// @func: freecacheTest
// @date: 2023-11-01 22:18:32
// @brief: 测试fresscache用法
// @author: Kewin Li
func freecacheTest() {
	cache := freecache.NewCache(totalCacheSize)

	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")
	expire := 10 // expire in 60 seconds
	cache.Set(key, val, expire)
	got, err := cache.Get(key)

	if err != nil {
		fmt.Println(err)
	} else {
		//for {
		//	got, _ = cache.Get(key)
		//	leftTime, err := cache.TTL(key)
		//	fmt.Printf("key:\"%v\" ,val: %v, ttl:%v [%v]\n", string(key), string(got), leftTime, err)
		//	time.Sleep(time.Second)
		//}
		fmt.Printf("%v \n", got)
	}
	affected := cache.Del(key)

	fmt.Println("deleted key ", affected)
	fmt.Println("entry count ", cache.EntryCount())
}

func main() {
	args := os.Args

	switch args[1][0] {
	case '1':
		freecacheTest()
	}
}
