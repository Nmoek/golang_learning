package main

import (
	"errors"
	"fmt"
	"os"
)

// @func: test4
// @brief: recover函数使用
// @author: Kewin Li
func test4() {

	// 有时需要提示致命错误，但不期望程序直接崩溃
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v \n", err)
		}
	}()

	arr := []int{1}
	arr[2] = 3
}

// @func: test3
// @brief: panic的使用
// @author: Kewin Li
func test3() {

	// 1. 显式调用panic
	// panic("显示调用panic!\n")

	// 2. 数组越界
	arr := make([]int, 0, 1)
	arr[1] = 0
	fmt.Printf("arr=%v \n", arr)

}

func m_div(a int, b int) (result float64, err error) {

	if b == 0 {
		return -1, errors.New("除数为0")
	}

	return float64(a / b), nil
}

// @func: test2
// @brief: error的应用场景
// @author: Kewin Li
func test2() {

	res, err := m_div(1, 2)
	fmt.Printf("res=%v, err=%v\n", res, err)

	res, err = m_div(1, 0)
	fmt.Printf("res=%v, err=%v\n", res, err)

}

// @func: test1
// @brief: error类型定义
// @author: Kewin Li
func test1() {
	err1 := fmt.Errorf("this is a err=%d!", 1)
	fmt.Printf("err1=%s \n", err1)

	err2 := errors.New("this is err2")
	fmt.Printf("err2=%s \n", err2)

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

	}

}
