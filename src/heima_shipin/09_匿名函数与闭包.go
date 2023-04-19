package main

import (
	"fmt"
)

func test1() {

	// 匿名函数基本用法
	// 注意: Go中任何东西定义了都要使用，否则编译不过
	func() {
		fmt.Printf("i'm 匿名函数\n")
	}()

	// 可以通过函数类型声明承接
	f1 := func() {
		fmt.Printf("i'm 匿名函数2  \n")
	}

	type MyFunc func()
	var m_f MyFunc = f1
	m_f()

	//可以传参+返回值
	ret := func(i int, j int) (res int) {
		fmt.Printf("i=%d, j=%d \n", i, j)
		return i + j
	}(1, 2)
	fmt.Printf("ret=%d \n", ret)

}

func main() {
	test1()
}
