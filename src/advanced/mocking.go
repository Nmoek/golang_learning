/*
 * @file: mocking.go
 * @brief: mocking模拟运行实例
 * @author: Kewin Li
 * @date:2023-04-02
 */

package mocking_test

import (
	"fmt"
	"io"
	"os"
	"time"
)

var finalWord string = "Go!"
var startNum int = 3

func Count(out io.Writer, n int) {

	for i := n; i > 0; i-- {
		// Go的计时间隔默认以ns为单位，也难怪并发控制精细
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	fmt.Fprintln(out, "Go!")
}

func main() {
	Count(os.Stdout, startNum)
}
