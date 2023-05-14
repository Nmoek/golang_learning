package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person //作为一个匿名字段存在
	//等价于以下的定义
	// name string
	// age  int
	// sex  byte
	id      string
	address string
}

type Student2 struct {
	Person //作为一个匿名字段存在
	//等价于以下的定义
	// name string
	// age  int
	// sex  byte
	name string //和匿名字段有冲突
}

type Student3 struct {
	Person //作为一个匿名字段存在
	//等价于以下的定义
	// name string
	// age  int
	// sex  byte
	int
	string
}

// @func: main
// @brief: 结构体中匿名字段
// @author: Kewin Li
func main() {

	// 1. 顺序初始化
	s1 := Student{Person{"liming", 34, 1}, "123456", "云南省"}

	fmt.Printf("s1=%v \n", s1)
	// %+v 打印更详细的信息
	fmt.Printf("s1=%+v \n", s1)

	fmt.Printf("--------------------------\n")

	// 2. 部分初始化
	s2 := Student{Person: Person{name: "wanghong", age: 45}, id: "456789"}
	fmt.Printf("s2=%+v \n", s2)

	fmt.Printf("--------------------------\n")

	// 3. 匿名字段直接访问即可
	s2.name = "huluan"
	s2.age = 66
	fmt.Printf("s2=%+v \n", s2)

	// 整体赋值也可
	s2.Person = Person{"hutinahua", 34, 1}
	fmt.Printf("s2=%+v \n", s2)

	fmt.Printf("--------------------------\n")

	// 4. 匿名字段出现冲突时默认采取"就近原则"，不会访问到匿名字段
	s3 := Student2{}
	s3.Person.name = "iner_test" //显示访问内部name
	s3.name = "test"             //访问的是外部name
	s3.age = 45
	s3.sex = 1
	fmt.Printf("s3=%+v \n", s3)

	fmt.Printf("--------------------------\n")

	// 5. 非结构体的匿名字段访问
	s4 := Student3{}
	s4.int = 6
	s4.string = "test3"
	fmt.Printf("s4=%+v \n", s4)

}
