package main

import (
	"fmt"
	"os"
	"strconv"
)

// @func: testParse
// @brief: Parse函数将字符串类型转换为其他类型
// @author: Kewin Li
func testParse() {
	val1, err1 := strconv.ParseInt("121561", 10, 64)
	if err1 == nil {
		fmt.Printf("val=%v \n", val1)
	} else {
		fmt.Printf("err=%v \n", err1.Error())
	}

	fmt.Printf("--------------------\n")

	val2, err2 := strconv.ParseFloat("65.615", 64)
	if err2 == nil {
		fmt.Printf("val=%v \n", val2)
	} else {
		fmt.Printf("err=%v \n", err2.Error())
	}

	fmt.Printf("--------------------\n")

	val4, err4 := strconv.ParseComplex("6+5i", 64)
	if err4 == nil {
		fmt.Printf("val=%v\n", val4)
	} else {
		fmt.Printf("err=%v \n", err4.Error())
	}
}

// @func: testFormat
// @brief: Format函数将其他类型转换为字符串类型
// @author: Kewin Li
func testFormat() {

	// 1. 其他类型----->字符串
	s1 := strconv.FormatInt(51615, 10)
	fmt.Printf("s1=%s \n", s1)
	fmt.Printf("--------------------\n")

	s2 := strconv.FormatBool(true)
	fmt.Printf("s2=%s \n", s2)

	fmt.Printf("--------------------\n")

	s3 := strconv.FormatFloat(3.14159, 'f', -1, 64)
	fmt.Printf("s3=%s \n", s3)

}

// @func: testAppend
// @brief: Append函数可以将多种类型转换为字符串后进行追加
// @author: Kewin Li
func testAppend() {

	// 1. 整型转换 支持进制指定
	s1 := make([]byte, 0, 4)
	s1 = strconv.AppendInt(s1, 12345, 10)
	fmt.Printf("s1=%s \n", s1)
	s1 = append(s1, '+')
	s1 = strconv.AppendInt(s1, 6789, 16)
	fmt.Printf("s1=%s \n", s1)
	s1 = append(s1, '+')
	s1 = strconv.AppendInt(s1, 1111, 8)
	fmt.Printf("s1=%s \n", s1)

	fmt.Printf("------------------------\n")

	// 2. bool型转换
	s2 := make([]byte, 0, 4)
	s2 = strconv.AppendBool(s2, true)
	fmt.Printf("s2=%s \n", s2)

	fmt.Printf("------------------------\n")

	// 3. 添加字符串时带双引号
	s3 := make([]byte, 0, 4)
	s3 = strconv.AppendQuote(s3, "66666")
	fmt.Printf("s3=%s \n", s3)

	// 4. 添加字符时带单引号
	s4 := make([]byte, 0, 4)
	s4 = strconv.AppendQuoteRune(s4, '2')
	fmt.Printf("s4=%s \n", s4)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		testAppend()
	case '2':
		testFormat()
	case '3':
		testParse()
	}

}
