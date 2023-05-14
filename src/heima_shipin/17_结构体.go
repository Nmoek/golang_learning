package main

import (
	"fmt"
	"os"
)

// Student @brief: 学生结构体
type Student struct {
	age  int
	name string
	sex  byte // 1: 男 0:女
}

// @func: test1
// @brief: 结构体定义、初始化以及基本使用
// @author: Kewin Li
func test1() {
	// 初始化时顺序初始化，不指定字段时必须全部初始化
	// var s1 Student = Student{1}
	s1 := Student{1, "liming", 1}
	fmt.Printf("s1=%v \n", s1)

	fmt.Printf("------------------------\n")

	// 指明字段可以部分初始化，未指明的处理为0/nil
	s2 := Student{age: 2, name: "wanghong"}
	fmt.Printf("s2=%v \n", s2)

	fmt.Printf("------------------------\n")

	//操作结构体成员
	s2.age = 3
	s2.name = "66666"
	s2.sex = 1

	fmt.Printf("s2=%v \n", s2)

	fmt.Printf("------------------------\n")

	//Go中成员访问也被简化，无论是实体还是指针都可以用'.'访问
	s3 := new(Student)
	s3.age = 4
	s3.name = "5615616"
	s3.sex = 0

	fmt.Printf("s3=%v \n", *s3)

}

func studentIsRquel(s1 Student, s2 Student) {
	if s1 == s2 {
		fmt.Printf("OK! \n")
	} else {
		fmt.Printf("NO!\n")
	}
}

// @func: test2
// @brief: 结构体的比较与赋值
// @author: Kewin Li
func test2() {

	s1 := Student{}
	s2 := Student{}

	studentIsRquel(s1, s2)
	fmt.Printf("------------------------\n")

	s1.age = 1
	s2.name = "liming"
	studentIsRquel(s1, s2)

	fmt.Printf("------------------------\n")

	s1 = s2
	studentIsRquel(s1, s2)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	}

}
