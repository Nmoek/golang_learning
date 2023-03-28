package main

import (
	"fmt"
)

/*
 *@brief: 所有字母对象的接口
 */
type Zimu interface {
	m_func() string
}

/*
 *@brief: 字母对象A
 */
type A struct {
}

/*
 *@brief: 字母对象A的某个方法
 */
func (a A) m_func() string {
	return "i'm A!!!"
}

/*
 *@brief: 字母对象B
 */
type B struct {
}

/*
 *@brief: 字母对象B的某个方法
 */
func (b B) m_func() string {
	return "i'm B!!!"
}

/*
 *@brief: 对外统一的抽象接口
 */
func get_zimu(z Zimu) string {
	return z.m_func()
}

func main() {
	a := A{}
	b := B{}
	fmt.Printf("%s \n", get_zimu(a))
	fmt.Printf("%s \n", get_zimu(b))
}
