package main

import (
	"fmt"
	"os"
)

// @func: test1
// @brief: append函数基本用法以及2倍扩容策略
// @author: Kewin Li
func test1() {

	s1 := []int{}

	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len=%d, cap=%d \n", len(s1), cap(s1))
	// output: len=3, cap=4  至少这里可以看出扩容是2倍策略

	fmt.Printf("----------------------------- \n")
	fmt.Printf("append扩容变化: \n")
	s2 := make([]int, 0, 1)
	oldCap := cap(s2)
	newCap := 0
	for i := 0; i < 10; i++ {
		s2 = append(s2, 1)
		newCap = cap(s2)
		if oldCap < newCap {
			fmt.Printf("%d -----> %d \n", oldCap, newCap)
			oldCap = newCap
		}

	}

}

// @func: test2
// @brief:  copy函数基本用法
// @author: Kewin Li
func test2() {

	// copy函数的特点只会对应相当长度的空间进行值赋值

	s1 := []int{1, 2}
	s2 := []int{6, 6, 6, 6, 6}

	fmt.Printf("s1=%v, p=%p\ns2=%v, p=%p\n", s1, &s1, s2, &s2)

	copy(s2, s1)

	fmt.Printf("s1=%v, p=%p\ns2=%v, p=%p\n", s1, &s1, s2, &s2)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	}

}
