package main

import (
	"reflect" //对标深拷贝-->深比较, 不做类型检查单纯字节比较
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("this is 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got=%d want=%d given=%v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got=%v want=%v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	errorMsg := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got=%v want=%v \n", got, want)
		}
	}

	//截取尾部切片合并为新切片
	t.Run("Sum All Tail with Slice", func(t *testing.T) {
		got := SumAllTailwithSlice([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		errorMsg(t, got, want)
	})

	//指定某一个元素合并为新切片
	t.Run("Sum All Tail with Num", func(t *testing.T) {
		got := SumAllTailwithNum([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		errorMsg(t, got, want)
	})

	//合并一个空切片
	t.Run("Sum an empty Slice", func(t *testing.T) {

		got := SumAllTailwithSlice([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		errorMsg(t, got, want)

	})

}

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
		tail := slice[1] //取出slice[1]这一个单独元素
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
