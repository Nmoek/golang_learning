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

// 计算每一个切片的和，把所有切片的和一次放到一个新的切片中返回
func SumAll(slices ...[]int) (sums []int) {
	slice_nums := len(slices)
	sums = make([]int, slice_nums) // 构建一个新的切片

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return
}

func Sum(numbers []int) int {
	var ret int
	for _, number := range numbers {
		ret += number
	}

	return ret
}
