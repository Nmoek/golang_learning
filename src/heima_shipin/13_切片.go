package main

import (
	"fmt"
	"os"
)

// @func: slice3
// @brief: 数组与切片的关系
// @author: Kewin Li
func slice3() {

	// 切片的修改会影响原数组
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr= %v \n", arr)

	s1 := arr[:3]
	s1[0] = 666
	fmt.Printf("s1=%p, %v\n", &s1, s1)
	fmt.Printf("arr=%p, %v \n", &arr, arr)

	fmt.Printf("---------------------- \n")

	// 原数组的修改也会影响切片
	arr2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr2= %v \n", arr2)
	s2 := arr2[3:]
	arr2[4] = 666
	fmt.Printf("s2=%p, %v \n", &s2, s2)
	fmt.Printf("arr2=%p, %v\n", &arr2, arr2)
	s2 = append(s2, 777)
	//s2超过容量上限，会在底层重新开辟新数组copy数据后，再进行追加
	// 因此，追加的内容没有影响到原数组
	fmt.Printf("append s2=%p, %v \n", &s2, s2)
	fmt.Printf("append arr2=%p, %v\n", &arr2, arr2)

	fmt.Printf("---------------------- \n")

	// ！！注意: 切片类似一种数组的引用,
	s3 := arr[1:4]
	fmt.Printf("s3=%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	s3 = append(s3, 777)
	fmt.Printf("append s3=%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("append arr=%p, %v\n", &arr, arr)

	fmt.Printf("---------------------- \n")

	//s3=	2 3 4
	// low=s3[2]=4, high=arr的长度上限
	s4 := s3[2:5]
	fmt.Printf("s4=%v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

}

// @func: slice2
// @brief: 切片创建
// @author: Kewin Li
func slice2() {

	// 1. 直接初始化
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s1=%v \n", s1)

	// 2. 使用make创建  make(type, len, cap)
	s2 := make([]int, 5, 5)
	fmt.Printf("s2=%v \n", s2)
}

// @func: slice1
// @brief: 切片截取以及长度与容量
// @author: Kewin Li
func slice1() {
	// [low:high:max]
	// low: 截取起点
	// high: 截取终点
	// max: 最大容量max-low

	arr := [...]int{1, 2, 3, 4, 5}
	s1 := arr[0:3:5]

	fmt.Printf("s1=%T, len=%d, cap=%d, %v \n", s1, len(s1), cap(s1), s1)
	// len = 3 - 0
	// cap = 5 - 0

	s2 := arr[1:4:4]
	fmt.Printf("s2=%T, len=%d, cap=%d, %v \n", s2, len(s2), cap(s2), s2)
	// len = 4 - 1
	// cap = 4 - 1

	// 等于 arr[0:3:5]
	s3 := arr[:3]
	fmt.Printf("s3=%T, len=%d, cap=%d, %v \n", s3, len(s3), cap(s3), s3)

	// 等于 arr[3:5:5]
	s4 := arr[3:]
	fmt.Printf("s4=%T, len=%d, cap=%d, %v \n", s4, len(s4), cap(s4), s4)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		slice1()
	case '2':
		slice2()
	case '3':
		slice3()
	}

}
