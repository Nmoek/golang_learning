package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var val1 int32 = 12

	// 加载值
	val2 := atomic.LoadInt32(&val1)
	fmt.Printf("val2=%d \n", val2)

	// 递增值
	val1 = atomic.AddInt32(&val1, 1)

	//CAS操作
	isValid := atomic.CompareAndSwapInt32(&val1, 13, 15)
	fmt.Printf("isValid:%v \n", isValid)
}
