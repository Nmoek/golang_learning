/*
 * @file:sum.go
 * @brief: 数组与切片slice练习
 * @author: Kewin Li
 * @date: 2023-03-31
 */

package main

import (
	"fmt"
	"os" //提供操作系统接口
)

// 每个切片的尾部切片合并到同一个切片中
func SumAllTailwithSlice(slices ...[]int) []int {

	var sums []int
	for _, slice := range slices {
		tail := slice[1:] //截取出从slice[1]开始一直到末尾的切片
		sums = append(sums, Sum(tail))
	}

	return sums
}

// 每个切片的尾部单个元素合并到同一个切片中
func SumAllTailwithNum(slices ...[]int) []int {

	var sums []int
	for _, slice := range slices {
		tail := slice[1]
		sums = append(sums, tail)
	}

	return sums
}

// 计算每一个切片的和，把所有切片的和依次放到一个新的切片中返回
func SumAll(slices ...[]int) (sums []int) {
	slice_nums := len(slices)
	sums = make([]int, slice_nums) // 构建一个新的切片

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return sums
}

func Sum(numbers []int) int {
	var ret int
	for _, number := range numbers {
		ret += number
	}

	return ret
}

// 带命令参数

func main() {

	//获取命令行参数
	argv := os.Args
	if len(argv) <= 1 {
		fmt.Printf("input param: go run xx.go 1\n")
		return
	}

	// fmt.Printf("argv=%v argv type= %T \n", argv, argv)

	cmd := argv[1][0]
	switch cmd {
	case '0':
		{
			got := Sum([]int{1, 2})
			want := 3
			fmt.Printf("got=%v want=%v \n", got, want)
		}

	case '1':
		{
			got := SumAll([]int{1, 2}, []int{1, 2})
			want := []int{3, 3}
			fmt.Printf("got=%v want=%v \n", got, want)
		}
	case '2':
		{
			got := SumAllTailwithNum([]int{1, 2}, []int{5, 6})
			want := []int{2, 6}
			fmt.Printf("got=%v want=%v \n", got, want)
		}
	case '3':
		{
			got := SumAllTailwithSlice([]int{1, 2}, []int{0, 9})
			want := []int{2, 9}
			fmt.Printf("got=%v want=%v \n", got, want)
		}
	default:
		{
			fmt.Printf("input cmd err! cmd=%c\n", cmd)
		}
	}

	return
}
