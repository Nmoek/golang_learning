package main

import (
	"28_generics/func_generics"
	"28_generics/struct_generics"
	"fmt"
	"os"
	"strconv"
)

// @func: sumTest
// @date: 2023年9月30日
// @brief: 切片求和测试
// @author: Kewin Li
func sumTest() {
	res1 := func_generics.Sum[int]([]int{1, 2, 3})

	fmt.Printf("int type res1=%v \n", res1)

	fmt.Printf("---------------\n")

	res2 := func_generics.Sum[float64]([]float64{6.3, 7.0, 4.5})

	fmt.Printf("float type res2=%v \n", res2)
}

// @func: maxTest
// @date: 2023年9月30日
// @brief: 切片最大值/最小值测试
// @author: Kewin Li
func max_minTest() {
	res1 := func_generics.Max[int]([]int{1, 2, 3})

	fmt.Printf("int type max res1=%v \n", res1)

	fmt.Printf("---------------\n")

	res2 := func_generics.Max[float64]([]float64{6.3, 7.0, 4.5})

	fmt.Printf("float type max res2=%.1f \n", res2)

	fmt.Printf("*****************\n")

	res1 = func_generics.Min[int]([]int{1, 2, 3})

	fmt.Printf("int type min res1=%v \n", res1)

	fmt.Printf("---------------\n")

	res2 = func_generics.Min[float64]([]float64{6.3, 7.0, 4.5})

	fmt.Printf("float type min res2=%.1f \n", res2)
}

// @func: findTest
// @date: 2023年9月30日
// @brief: 切片过滤与查找测试
// @author: Kewin Li
func findTest() {

	// 1. 数字类测试
	res1, f1 := func_generics.Find[int]([]int{1, 2, 3, 11, 300}, func(v int) bool {
		return v > 5
	})

	if f1 {
		fmt.Printf("int type find res1=%v \n", res1)
	} else {
		fmt.Printf("int type not find \n")
	}

	fmt.Printf("---------------\n")

	res2, f2 := func_generics.Find[float64]([]float64{0.2, 1.2, 1.0}, func(v float64) bool {
		return v > 2.0
	})

	if f2 {
		fmt.Printf("float type find res2=%v \n", res2)
	} else {
		fmt.Printf("float64 type not find \n")
	}

	fmt.Printf("*****************\n")

	// 2. 字符串类测试
	res3, f3 := func_generics.Find[string]([]string{"liming", "wanggang", "12345"}, func(str string) bool {
		// 检查字符串中是否有数字
		_, b := strconv.Atoi(str)

		if b != nil {
			return true
		}
		return false
	})

	if f3 {
		fmt.Printf("string type find res3=%v \n", res3)
	} else {
		fmt.Printf("string type not find \n")
	}

}

// @func: insertTest
// @date: 2023年9月30日
// @brief: 切片指定位置插入元素测试
// @author: Kewin Li
func insertTest() {
	res1 := func_generics.Insert[int](0, 100, []int{1, 2, 3})

	if res1 != nil {
		fmt.Printf("int type insert res1=%v \n", res1)
	}

	fmt.Printf("---------------\n")

	res2 := func_generics.Insert[float64](2, 30.4, []float64{6.3, 7.0, 4.5})

	if res2 != nil {
		fmt.Printf("float type insert res2=%.1f \n", res2)
	}

	fmt.Printf("---------------\n")

	res3 := func_generics.Insert[string](3, "ljk", []string{"111", "222", "333"})

	if res3 != nil {
		fmt.Printf("string type insert res3=%v \n", res3)
	}
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("input xxx 2 \n")
		return
	}

	switch args[1][0] {
	case '1':
		struct_generics.UseList()
	case '2':
		sumTest()
	case '3':
		max_minTest()
	case '4':
		findTest()
	case '5':
		insertTest()
	}

}
