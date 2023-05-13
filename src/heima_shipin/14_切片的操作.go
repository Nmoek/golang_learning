package main

import (
	"fmt"
)

// @func: test1
// @brief: append函数基本用法
// @author: Kewin Li
func test1() {

	s1 := []int{}

	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	// output: len=3, cap=4  至少这里可以看出扩容是2倍策略
}

// @func: test2
// @brief: append函数的2倍扩容策略
// @author: Kewin Li
func test2() {

	s1 := make([]int, 0, 1)

	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	//output: old len=0, cap=1

	s1 = append(s1, 1)
	s1 = append(s1, 1)

	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	//output: len=2, cap=2

	s1 = append(s1, 2)
	s1 = append(s1, 2)
	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	//output: len=4, cap=4

	s1 = append(s1, 3)
	s1 = append(s1, 3)
	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	//output: len=6, cap=8
}

func main() {

	// test1()

	test2()
}
