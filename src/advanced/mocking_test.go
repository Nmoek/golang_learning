/*
 * @file: mocking_test.go
 * @brief: 模拟mocking测试
 * @author: Kewin Li
 * @date: 2023-04-03
 */

package mocking_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// Sleeper @brief: 自定义睡眠接口
type Sleeper interface {
	Sleep()
}

type SpySleep struct {
	Calls int
}

// @func: Sleep
// @brief: 模拟睡眠(实际没有睡眠，而是将睡眠动作转换为计数)
// @author: Kewin Li
// @receiver: *SpySleep s
func (s *SpySleep) Sleep() {
	s.Calls++
}

func Count(out io.Writer, n int, s *SpySleep) {
	for i := n; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}

	s.Sleep()
	fmt.Fprintln(out, "GO!")
}

func TestCount(t *testing.T) {

	// 输出到标准输出
	t.Run("out stdout", func(t *testing.T) {
		spy := &SpySleep{}
		Count(os.Stdout, 3, spy)
	})

	//输出到指定buffer
	t.Run("out buffer", func(t *testing.T) {

		buffer := bytes.Buffer{}

		spy := &SpySleep{}

		Count(&buffer, 3, spy)

		got := buffer.String()
		want := `3
2
1
GO!
`
		if got != want {
			t.Errorf("got= '%s' want='%s' \n", got, want)
		}

		if spy.Calls != 4 {
			t.Errorf("not enough wait seconds!\n")
		}

	})

}
