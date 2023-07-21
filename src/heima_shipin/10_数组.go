package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func test2(nums *[10]int) {
	(*nums)[0] = 666
}

func test1(nums [10]int) {
	nums[0] = 666
}

// @func: myFindMax(helper function)
// @brief: 找出数组中最大值的下标以及数值
// @date: 2023年7月17日
// @author: Kewin Li
// @param: [10]int nums
// @return idx
// @return valMax
func myFindMax(nums [10]int) (idx int, valMax int) {

	valMax = -1

	for i, v := range nums {
		if v > valMax {
			idx = i
			valMax = v
		}
	}

	return idx, valMax
}

// @func: numbers7
// @brief: 多维数组
// @date: 2023年7月17日
// @author: Kewin Li
func numbers7() {

	rand.Seed(time.Now().UnixMicro())

	var nums2 [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			nums2[i][j] = rand.Intn(100)
		}
	}

	fmt.Printf("nums2=%v, addr0=%p, addr1=%p, len=%d, cap=%d\n", nums2, &nums2[0], &nums2[1], len(nums2), cap(nums2))

	fmt.Printf("#####################\n")

	// 1. 外层数组切片引用
	slice1 := nums2[1:]

	fmt.Printf("before slice1=%v, addr=%p, len=%d, cap=%d\n", slice1, &slice1, len(slice1), cap(slice1))

	tmp := [2]int{666}
	slice1 = append(slice1, tmp)
	fmt.Printf("nums2=%v, addr0=%p, addr1=%p, len=%d, cap=%d\n", nums2, &nums2[0], &nums2[1], len(nums2), cap(nums2))

	fmt.Printf("after slice1=%v, addr=%p, len=%d, cap=%d\n", slice1, &slice1, len(slice1), cap(slice1))

	fmt.Printf("--------------------------\n")

	// 2. 内层数组切片引用
	slice2 := nums2[1][1:]

	fmt.Printf("before slice2=%v, addr=%p, len=%d, cap=%d\n", slice2, &slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 999)
	fmt.Printf("nums2=%v, addr0=%p, addr1=%p, len=%d, cap=%d\n", nums2, &nums2[0], &nums2[1], len(nums2), cap(nums2))

	fmt.Printf("after slice2=%v, addr=%p, len=%d, cap=%d\n", slice2, &slice2, len(slice2), cap(slice2))

}

// @func: numbers6
// @brief: 26字母遍历
// @date: 2023年7月17日
// @author: Kewin Li
func numbers6() {

	zimu := [26]byte{0: 'A'}

	for i := 0; i < 26; i++ {
		zimu[i] = byte('A' + i)
	}

	for _, v := range zimu {
		fmt.Printf("%c ", v)
	}

	fmt.Printf("\n")

}

func numbers5() {

	var nums [5]float64
	fmt.Printf("input float:\n")
	for i := 0; i < len(nums); i++ {
		fmt.Scanf("%f", &nums[i])
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%.2f ", nums[i])
	}

	fmt.Printf("\n")

}

func numbers4() {

	rand.Seed(time.Now().UnixMicro())
	var nums [10]int
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(100)
	}

	test1(nums)
	fmt.Printf("nums=%v \n", nums)

	test2(&nums)
	fmt.Printf("nums=%v \n", nums)
}

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
	pa0 := &a[0]
	pa1 := &a[1]
	fmt.Printf("a=%T:%v, len=%d \n", a, a, len(a))

	fmt.Printf("b=%T:%v, len=%d \n", b, b, len(b))

	fmt.Printf("pa=%T: %p \n", pa, pa)

	fmt.Printf("pa0=%T: %p \n", pa0, pa0)

	fmt.Printf("pa1=%T: %p \n", pa1, pa1)

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

	case '4':
		numbers4()

	case '5':
		numbers5()

	case '6':
		numbers6()

	case '7':
		numbers7()
	}

}
