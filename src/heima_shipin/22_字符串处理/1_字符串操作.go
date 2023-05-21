package main

import (
	"fmt"
	"os"
	"strings"
)

// @func: testFields
// @brief: 将字符串s的空格去除，并按空格分割返回切片
// @author: Kewin Li
func testFields() {

	s1 := "1 2 3 4 5 6"

	fmt.Printf("s1=%s, res=%s \n", s1, strings.Fields(s1))

}

// @func: testTrim
// @brief: 将字符串s头部尾部指定的字符(串)dep全部去除
// @author: Kewin Li
func testTrim() {

	// func Trim(s, cutset string) string

	s1 := "!!!!!12354!!!!!!"

	fmt.Printf("s1=%s, res=%s \n", s1, strings.Trim(s1, "!"))

}

// @func: testSplit
// @brief: 将字符串s按dep串分割为切片
// @author: Kewin Li
func testSplit() {

	s1 := "1,2,3,4,5"
	fmt.Printf("s1=%s, res=%v \n", s1, strings.Split(s1, ","))

	fmt.Printf("s1=%s, res=%v \n", s1, strings.Split(s1, "2"))

}

// @func: testReplace
// @brief: 字符串s中指定的子串old替换为新串new, 指定替换n个, -1全换
// @author: Kewin Li
func testReplace() {

	// func Replace(s, old, new string, n int) string

	s1 := "1i3 1i3 1i3"
	fmt.Printf("s1=%s, res=%s \n", s1, strings.Replace(s1, "i", "2", 4))
	fmt.Printf("s1=%s, res=%s \n", s1, strings.Replace(s1, "i", "2", 1))
	fmt.Printf("s1=%s, res=%s \n", s1, strings.Replace(s1, "i", "2", -1))

}

// @func: testRepeat
// @brief: 将字符串s复制拼接在一起返回
// @author: Kewin Li
func testRepeat() {

	// func Repeat(s string, count int) string

	// 1. 基本用法
	s1 := "123"
	fmt.Printf("s1=%s, res=%s \n", s1, strings.Repeat(s1, 3))

	fmt.Printf("------------\n")

	// 2. 复制0次则返回空字符串
	s2 := "1111"

	fmt.Printf("s2=%s, res=%s \n", s2, strings.Repeat(s2, 0))

	fmt.Printf("------------\n")

	// 3. 负数会panic
	s3 := "159619"
	s3 = strings.Repeat(s3, -1)

}

// @func: testIndex
// @brief: 返回子串substr在字符串s中首字母的下标位置
// @author: Kewin Li
func testIndex() {

	// func Index(s, substr string) int

	s := "123456"

	i := strings.Index(s, "234")
	fmt.Printf("i=%d \n", i)

	fmt.Printf("--------------\n")

	i = strings.Index(s, "adaf")
	fmt.Printf("i=%d \n", i) //找不到结果就返回-1

}

// @func: testJion
// @brief: 将多个字符串[]string, 以dep字符串连接组成新字符串
// @author: Kewin Li
func testJoin() {

	// func Join(elems []string, sep string) string

	s1 := []string{"1", "2", "3", "4", "5"}

	s2 := strings.Join(s1, ",")
	fmt.Printf("s2=%v \n", s2)

	s2 = strings.Join(s1, "")
	fmt.Printf("s2=%v \n", s2)

}

// @func: testContains
// @brief: 字符串s中是否包含子串substr
// @author: Kewin Li
func testContains() {

	// func Contains(s, substr string) bool

	// 1. s和substr都不为空
	s := "123456789"
	f := strings.Contains(s, string("3456"))

	fmt.Printf("f=%v \n", f)

	fmt.Printf("--------------------\n")

	// 2. s不空，substr空
	f = strings.Contains(s, "")
	fmt.Printf("f=%v \n", f)

	fmt.Printf("--------------------\n")

	// 3. s空，substr不空
	f = strings.Contains("", "123456")
	fmt.Printf("f=%v \n", f)

	fmt.Printf("--------------------\n")

	// 4. substr长度超过s
	f = strings.Contains(s, "1234567890")
	fmt.Printf("f=%v \n", f)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		testContains()
	case '2':
		testJoin()
	case '3':
		testIndex()
	case '4':
		testRepeat()
	case '5':
		testReplace()
	case '6':
		testSplit()
	case '7':
		testTrim()
	case '8':
		testFields()
	}

}
