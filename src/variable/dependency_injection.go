/*
 * @file: dependency_injection_test.go
 * @brief: 依赖注入测试学习
 * @author: Kewin Li
 * @date: 2023-04-02
 */

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello %s", name)
}

func main() {

	// 向标准输出中写入数据
	Greet(os.Stdout, "ljk \n")

	// 向自定义的buffer中写入数据
	buffer := bytes.Buffer{}
	Greet(&buffer, "ljk2")
	fmt.Printf("buffer= '%s' \n", buffer.String())

}
