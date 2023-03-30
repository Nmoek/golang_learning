/*
 * @file: iteration_test.go
 * @brief: 引入用例测试与基准测试
 * @author: Kewin Li
 * @date:2023-03-31
 */

package main

import (
	"fmt"
	"testing"
)

/*
 * 基准测试(benchmarks)
 * 命令: $go test -bench=.
 */
func BenchmarkReapt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

// 用例测试
func ExampleRepeat() {
	repeat := Repeat("b", 5)
	fmt.Printf("%s\n", repeat)
	// Output: bbbbb
}

// 常规测试
func TestRepeat(t *testing.T) {

	repeat := Repeat("a", 5)
	expect := "aaaaa"

	if repeat != expect {
		t.Errorf("repeat='%q' expect='%q' \n", repeat, expect)
	}

}

func Repeat(str string, count int) string {

	var ret string
	for i := 0; i < count; i++ {
		ret += str
	}

	return ret
}
