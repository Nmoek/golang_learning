package main

import (
	"fmt"
	"os"
)

// @func: numbers3
// @brief: 数组比较与相互赋值
// @author: Kewin Li
func numbers3() {

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 4}

	fmt.Printf("a == b: %v \n", a == b)
	fmt.Printf("a == c: %v \n", a == c)

	// err: 同一种类型的数组才能相互比较与赋值
	// d := [4]int{1, 2, 3, 4}
	// fmt.Printf("a == d: %v \n", a == d)

}

// @func: numbers2
// @brief: 二维数组的本质
// @author: Kewin Li
func numbers2() {

	a := [2][2]int{{1, 2}, {3, 4}}
	b := a[0]
	pa := &a
	pb := &a[1]
	fmt.Printf("a=%T:%v, len=%d \n", a, a, len(a))

	fmt.Printf("b=%T:%v, len=%d \n", b, b, len(b))

	fmt.Printf("pa=%T: %p \n", pa, pa)

	fmt.Printf("pb=%T: %p \n", pb, pb)
}

// @func: numbers1
// @brief: 数组初始化
// @author: Kewin Li
func numbers1() {

	//1. 全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	//a := [5]int{1,2,3,4,5}
	fmt.Printf("a=%T:%v len=%d \n", a, a, len(a))

	// 2. 部分初始化
	b := [5]int{1, 2, 3}
	fmt.Printf("b=%T:%v, len=%d \n", b, b, len(b))

	// 3.指定位置初始化
	c := [5]int{2: 10, 4: 6}
	fmt.Printf("c=%T:%v, len=%d\n", c, c, len(c))
}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		numbers1()

	case '2':
		numbers2()

	case '3':
		numbers3()
	}

}
