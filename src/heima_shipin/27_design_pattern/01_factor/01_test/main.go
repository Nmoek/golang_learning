package main

import (
	"01_test/model"
	"fmt"
)

func main() {

	// 非法
	// stu1 := model.student{"xiaoming", 23, 64.7}

	// 工厂模式解决
	stu1 := model.NewStudent("xiaoming", 23, 64.7)

	// 1. 当struct中的字段首字母大写外部可以进行访问
	// type student struct {
	// 	Name   string
	// 	Age    int
	// 	Fenshu float64
	// }

	// fmt.Printf("stu1: name=%s, age=%d, fenshu=%.2f \n", stu1.Name, stu1.Age, stu1.Fenshu)

	//  2. 当struct中的字段首字母小写，外部无法访问，需要借助接口
	// type student struct {
	// 	name   string
	// 	age    int
	// 	fenshu float64
	// }
	fmt.Printf("stu1: name=%s, age=%d, fenshu=%.2f \n", model.GetName(*stu1), model.GetAge(*stu1), model.GetFenshu(*stu1))

	// 3.无论struct中的字段首字母大写或小写，都能直接%v打印
	// fmt.Printf("stu1=%+v \n", *stu1)
}
