package main

import (
	"fmt"
	"os"
)

// long @brief: 自定义类型
type long int

// person @brief:  person结构体
type person struct {
	name string
	age  int
	sex  byte
}

// student @brief: student结构体
type worker struct {
	person  //不光继承字段，对应的方法将一并继承
	address string
}

// @func: Add1
// @brief: 面向过程的Add函数
// @author: Kewin Li
// @param: int a
// @param: int b
// @return int
func Add1(a int, b int) int {
	return a + b
}

// @func: Add2
// @brief: 面向对象的Add函数（long类型的方法）
// @author: Kewin Li
// @receiver: long l
// @param: long n
// @return long
func (l long) Add2(n long) long {
	return l + n
}

// @func: printfInfo
// @brief: 打印person结构体中所有变量信息
// @author: Kewin Li
// @receiver: person p
func (p person) printfInfo() {

	fmt.Printf("%+v \n", p)
}

// @func: setName
// @brief: 给实例修改Name
// @author: Kewin Li
// @receiver: person p
// @param: string name
func (p *person) setName(name string) {
	p.name = name
}

// @func: setAge
// @brief: 给实例修改Age
// @author: Kewin Li
// @receiver: person p
// @param: int age
func (p *person) setAge(age int) {
	p.age = age
}

// @func: setSex
// @brief: 给实例修改Sex
// @author: Kewin Li
// @receiver: person p
// @param: byte sex
func (p *person) setSex(sex byte) {
	p.sex = sex
}

// @func: printfInfo
// @brief: worker结构体也实现和person结构体一样的方法
// @author: Kewin Li
// @receiver: worker w
func (w worker) printfInfo() {
	fmt.Printf("%+v\n", w)
}

// @func: test4
// @brief: 方法值以及方法表达式
// @author: Kewin Li
func test4() {

	// 1. 方法值
	w1 := worker{}
	// 有点儿类似C++的函数对象
	pFunc := w1.printfInfo
	pFunc() //可以隐藏接收者
	fmt.Printf("pFunc=%T \n", pFunc)

	fmt.Printf("-----------------------\n")

	// 2. 方法表达式
	w2 := worker{}
	vFunc := (*worker).setAge
	vFunc(&w2, 34)
	fmt.Printf("vFunc=%T \n", vFunc)

	vFunc2 := (worker).printfInfo
	vFunc2(w2)
	fmt.Printf("vFunc2=%T \n", vFunc2)
}

// @func: test3
// @brief: 方法的继承与重写
// @author: Kewin Li
func test3() {
	// 1. 匿名结构体的方法会被继承
	w1 := worker{person{"wanghong", 14, 1}, "浙江省"}
	w1.printfInfo()
	//注意: worker结构体中没有printfInfo()方法时，会调用匿名结构体的方法

	fmt.Printf("------------------------\n")

	w1.person.printfInfo()
}

// @func: test2
// @brief: 结构体添加方法
// @author: Kewin Li
func test2() {

	p1 := person{}
	p1.printfInfo()

	fmt.Printf("--------------------\n")
	p1.setAge(34)
	p1.setName("ljk")
	p1.setSex(1)
	p1.printfInfo()

}

// @func: test1
// @brief: 自定义类型添加方法
// @author: Kewin Li
func test1() {

	// 1. 面向过程编程
	res1 := Add1(1, 5)
	fmt.Printf("res1=%d \n", res1)

	// 2. 自定义类型添加方法
	var l long = 1
	res2 := l.Add2(2)
	fmt.Printf("res2=%d \n", res2)

}

type MethodUtils struct {
}

func (m MethodUtils) juzhen_10_8() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func (m MethodUtils) juzhen_m_n(w int, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func (m MethodUtils) juzhen_area(w, h int) float64 {

	return float64(w * h)
}

// @func: timu1
// @date: 2023年8月23日
// @brief: 通过方法打印10*8矩阵
// @author: Kewin Li
func timu1() {

	m1 := MethodUtils{}

	m1.juzhen_10_8()
}

// @func: timu2
// @date: 2023年8月23日
// @brief: 通过方法打印m*n矩阵
// @author: Kewin Li
func timu2() {
	m1 := MethodUtils{}

	m1.juzhen_m_n(5, 6)
}

// @func: timu3
// @date: 2023年8月23日
// @brief: 通过方法打印计算矩阵面积
// @author: Kewin Li
func timu3() {
	m1 := MethodUtils{}

	area1 := m1.juzhen_area(5, 6)

	fmt.Printf("area1= %f \n", area1)
}

type m_dog struct {
	name   string
	age    int
	weight float64
}

func (d m_dog) say() {

	fmt.Printf("m_dog info: %+v \n", d)
}

// @func: timu4
// @date: 2023年9月5日
// @brief: 定义一个狗对象，并实现打印方法
// @author: Kewin Li
func timu4() {

	dog1 := m_dog{"wangcai", 2, 23.5}

	dog1.say()

}

func deferReturnVal() {

	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("%d \n", i)
		}()
	}
}

func deferReturnVal2() {

	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("%d \n", val)
		}(i)
	}
}

func deferReturnVal3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("%d \n", j)
		}()
	}
}

// @func: defer_test
// @date: 2023年9月25日
// @brief: defer与闭包的使用
// @author: Kewin Li
func defer_test() {
	deferReturnVal()
	fmt.Printf("-------------\n")
	deferReturnVal2()
	fmt.Printf("-------------\n")
	// ??
	deferReturnVal3()
}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	case '3':
		test3()
	case '4':
		test4()
	// 课堂练习1
	case '5':
		timu1()
	// 课堂练习2
	case '6':
		timu2()
	// 课堂练习3
	case '7':
		timu3()
	case '8':
		timu4()
	case '9':
		defer_test()
	}

}
