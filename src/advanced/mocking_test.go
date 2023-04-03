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
	"time"
)

func Count(out io.Writer, n int) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprintln(out, "GO!")
}

func TestCount(t *testing.T) {

	// 输出到标准输出
	t.Run("out stdout", func(t *testing.T) {
		Count(os.Stdout, 3)
	})

	//输出到指定buffer
	t.Run("out buffer", func(t *testing.T) {

		buffer := bytes.Buffer{}

		Count(&buffer, 3)

		got := buffer.String()
		want := `3
2
1
GO!
`
		if got != want {
			t.Errorf("got= '%s' want='%s' \n", got, want)
		}

	})

}
