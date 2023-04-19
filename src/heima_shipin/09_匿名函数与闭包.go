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

func test2() {

	a := 10
	str := "ljk"

	//闭包的变量捕获是引用型的
	func() {
		fmt.Printf("a=%d, str=%s \n", a, str)
		a = 1
		str = "666"
	}()

	fmt.Printf("a=%d, str=%s \n", a, str)

}

func test3() int {
	var x int
	x++
	return x * x
}

func test4() func() (int, *int) {
	var x int
	return func() (int, *int) {
		x++
		p := &x
		return x * x, p
	}

}

func main() {
	test1()
	fmt.Printf("--------------------- \n")
	test2()
	fmt.Printf("--------------------- \n")
	fmt.Printf("闭包经典例子: \n")
	fmt.Printf("test3 res=%d \n", test3())
	fmt.Printf("test3 res=%d \n", test3())
	fmt.Printf("test3 res=%d \n", test3())
	f := test4() //注意这里返回的是闭包, f()多次调用，使用了同一个闭包多次
	f2 := test4()
	// _, pf := f()
	// _, pf2 := f2()
	fmt.Printf("f type=%T addr=%p\n", f, f)
	fmt.Printf("f2 type=%T addr=%p\n", f2, f2)

	fmt.Printf("--------------------- \n")
	//err: test4()()这种写法确实调用了闭包，但没有循环使用，
	// 连续生成了多个执行逻辑一样但内存分配不一样的闭包
	x1, p1 := test4()()
	x2, p2 := test4()()
	x3, p3 := test4()()
	x4, p4 := f()
	x5, p5 := f()
	x6, p6 := f()
	fmt.Println("test4 res=", x1, ",", p1)
	fmt.Println("test4 res= ", x2, ",", p2)
	fmt.Println("test4 res=", x3, ",", p3)
	fmt.Println("err test4 res=", x4, ",", p4)
	fmt.Println("err test4 res=", x5, ",", p5)
	fmt.Println("err test4 res=", x6, ",", p6)

}
