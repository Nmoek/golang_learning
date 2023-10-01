package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// student @brief: 学生结构体
type student struct {
	age  int
	name string
	sex  byte // 1: 男 0:女
}

// @func: test1
// @brief: 结构体定义、初始化以及基本使用
// @author: Kewin Li
func test1() {
	// 初始化时顺序初始化，不指定字段时必须全部初始化
	// var s1 student = student{1}
	s1 := student{1, "liming", 1}
	fmt.Printf("s1=%v \n", s1)

	fmt.Printf("------------------------\n")

	// 指明字段可以部分初始化，未指明的处理为0/nil
	s2 := student{age: 2, name: "wanghong"}
	fmt.Printf("s2=%v \n", s2)

	fmt.Printf("------------------------\n")

	//操作结构体成员
	s2.age = 3
	s2.name = "66666"
	s2.sex = 1

	fmt.Printf("s2=%v \n", s2)

	fmt.Printf("------------------------\n")

	//Go中成员访问也被简化，无论是实体还是指针都可以用'.'访问
	s3 := new(student)
	s3.age = 4
	s3.name = "5615616"
	s3.sex = 0

	fmt.Printf("s3=%v \n", *s3)

}

func studentIsRquel(s1 student, s2 student) {
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

	s1 := student{}
	s2 := student{}

	studentIsRquel(s1, s2)
	fmt.Printf("------------------------\n")

	s1.age = 1
	s2.name = "liming"
	studentIsRquel(s1, s2)

	fmt.Printf("------------------------\n")

	s1 = s2
	studentIsRquel(s1, s2)

}

// @func: test3
// @date: 2023年8月10日
// @brief: 检验结构体是否属于值类型
// @author: Kewin Li
func test3() {

	s1 := student{15, "ljk", 1}

	fmt.Printf("%+v \n", s1)

	s2 := s1

	s2.age = 22

	// 修改s2的值观察s1的值是否也改变
	fmt.Printf("%+v \n", s1)
	fmt.Printf("%+v \n", s2)

}

type m_struct struct {
	Age  int            `json:"age"`
	Name []byte         `json:"name"`
	Mp   map[int]string `json:"mp"`
}

// @func: test4
// @date: 2023年8月13日
// @brief: 证明结构体的值类型性质
// @author: Kewin Li
func test4() {

	m_st1 := m_struct{
		1,
		[]byte("ljk"),
		map[int]string{111: "111"},
	}

	fmt.Println("m_st1:", m_st1)

	m_st2 := m_st1

	m_st2.Name = []byte("zqc")
	m_st2.Mp[111] = "222"

	fmt.Println("m_st1:", m_st1)
	fmt.Println("m_st2:", m_st2)

	fmt.Printf("---------------------\n")

	fmt.Println("m_st1:", m_st1)
	m_pst1 := &m_st1

	m_pst1.Name = []byte("ljk666")

	fmt.Println("m_st1:", m_st1)

	fmt.Printf("---------------------\n")

	buf, _ := json.Marshal(m_st1)

	fmt.Printf("%v \n", string(buf))

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	case '3':
		test3()
	case '4':
		test4()

	}

}
