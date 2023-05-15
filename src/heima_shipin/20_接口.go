package main

import (
	"fmt"
	"os"
)

// Humaner @brief: 人类接口(interface)
type Humaner interface {
	sayHi()
}

type student struct {
	name string
	age  int
}

// @func: sayHi
// @brief: student实现sayHi()
// @author: Kewin Li
// @receiver: student s
func (s student) sayHi() {
	fmt.Printf("student info: %+v \n", s)
}

type teacher struct {
	address string
	group   int
}

// @func: sayHi
// @brief: teacher实现sayHi()
// @author: Kewin Li
// @receiver: teacher t
func (t teacher) sayHi() {
	fmt.Printf("teacher info: %+v \n", t)
}

type myStr string

// @func: sayHi
// @brief: myStr实现sayHi()
// @author: Kewin Li
// @receiver: myStr str
func (str myStr) sayHi() {
	fmt.Printf("myStr:%s \n", str)
}

// @func: whoSayHi
// @brief: 典型多态接口, 将抽象类型作为入参
// @author: Kewin Li
// @param: Humaner h
func whoSayHi(h Humaner) {
	h.sayHi()
}

// @func: test1
// @brief:
// @author: Kewin Li
func test1() {

	// 1. 接口定义、实现以及基本用法
	// h Humaner(interface)
	// |
	// |——	s student
	// |—— t teacher
	// |—— m_str myStr
	var h Humaner
	s := student{"liming", 34}
	h = s
	h.sayHi()

	t := teacher{"浙江省", 1}
	h = t
	h.sayHi()

	m_str := myStr("666666666666")
	h = m_str
	h.sayHi()

	fmt.Printf("------------------------\n")

	// 2. 多态表现1: 函数入参
	whoSayHi(s)
	whoSayHi(t)
	whoSayHi(m_str)

	fmt.Printf("------------------------\n")

	// 3. 多态表现2: 容器存储
	hs := make([]Humaner, 3)
	hs[0] = s
	hs[1] = t
	hs[2] = m_str

	for _, f := range hs {
		f.sayHi()
	}

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()

	}

}
