package main

import (
	"fmt"
	"os"
	"strings"
)

// @func: slice5
// @date: 2023年7月20日
// @brief: string与slice
// @author: Kewin Li
func slice5() {

	var str string = "hello world!"

	idx := strings.Index(str, "world")
	slice1 := str[idx:]

	fmt.Printf("str addr=%p, len=%d\n", &str, len(str))
	fmt.Printf("slice1 addr=%p, len=%d\n", &slice1, len(slice1))

	slice2 := append([]byte(slice1), '6')
	fmt.Printf("str addr=%p, len=%d\n", &str, len(str))
	fmt.Printf("slice2 addr=%p, len=%d\n", &slice2, len(slice2))
}

// @func: slice4
// @date: 2023年7月20日
// @brief: 切片slice的copy操作
// @author: Kewin Li
func slice4() {

	// 1. 目的切片的空间足够
	slice1 := []int{1, 2, 3}

	slice2 := make([]int, 10)
	fmt.Printf("slice2=%v \n", slice2)

	copy(slice2, slice1)
	fmt.Printf("slice2=%v \n", slice2)

	fmt.Printf("--------------------\n")

	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice3 addr=%p, len=%d, cap=%d \n", &slice3, len(slice3), cap(slice3))
	slice4 := make([]int, 3)
	fmt.Printf("before slice4 addr=%p, len=%d, cap=%d \n", &slice4, len(slice4), cap(slice4))

	copy(slice4, slice3)

	fmt.Printf("slice4=%v \n", slice4)

	fmt.Printf("after slice4 addr=%p, len=%d, cap=%d \n", &slice4, len(slice4), cap(slice4))

}

// @func: slice3
// @brief: 数组与切片的关系、append扩容机制
// @author: Kewin Li
func slice3() {

	// 切片的修改会影响原数组
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr= %v \n", arr)

	s1 := arr[:3] // 1 2 3
	s1[0] = 666   // 666 2 3
	fmt.Printf("s1=%p, %v\n", &s1, s1)
	fmt.Printf("arr=%p, %v \n", &arr, arr)
	// arr=xxx 666 2 3 4 5 6
	fmt.Printf("---------------------- \n")

	// 原数组的修改也会影响切片
	arr2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr2= %v \n", arr2)
	s2 := arr2[3:]                      // 4 5 6
	arr2[4] = 666                       // 1 2 3 4 666 6
	fmt.Printf("s2=%p, %v \n", &s2, s2) // s2=0x111 4 666 6
	fmt.Printf("arr2=%p, %v\n", &arr2, arr2)
	s2 = append(s2, 777)
	//s2超过容量上限，会在底层重新开辟新数组copy数据后，再进行追加
	// 因此，追加的内容没有影响到原数组
	fmt.Printf("append s2=%p, %v \n", &s2, s2)      // s2=0x222 4 5 6 777
	fmt.Printf("append arr2=%p, %v\n", &arr2, arr2) // aar=0x000 1 2 3 4 666 6

	fmt.Printf("---------------------- \n")

	// ！！注意: 切片类似一种数组的引用,
	s3 := arr[1:4] // 2 3 4
	fmt.Printf("s3=%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	s3 = append(s3, 777)
	fmt.Printf("append s3=%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("append arr=%p, %v\n", &arr, arr)

	fmt.Printf("---------------------- \n")

	//s3=2 3 4
	// low=s3[2]=4, high=arr的长度上限
	s4 := s3[2:5] // 4 777 6
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
	case '4':
		slice4()
	case '5':
		slice5()
	}

}
