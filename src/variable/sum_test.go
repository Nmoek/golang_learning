package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	number := [5]int{1, 2, 3, 4, 5}

	got := Sum(number)
	want := 15

	if got != want {
		// %v是万能输出
		t.Errorf("sum=%d want=%d given=%v\n", got, want, number)
	}

}

func Sum(number [5]int) int {
	var ret int
	for i := 0; i < 5; i++ {
		ret += number[i]
	}

	return ret
}
