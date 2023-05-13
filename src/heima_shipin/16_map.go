package main

import (
	"fmt"
)

// @func: test1
// @brief: map的基本用法:创建、赋值、遍历、删除
// @author: Kewin Li
func test1() {

	//1. 创建、赋值
	m1 := map[int]string{}
	// m1 := make(map[int]string, 1)
	m1[1] = "111"
	m1[2] = "222"
	m1[3] = "333"

	fmt.Printf("m1=%v, len=%d\n", m1, len(m1))

	fmt.Printf("-------------\n")
	//2. 遍历/返回值

	for key, val := range m1 {
		fmt.Printf("key=%v, val=%v \n", key, val)
	}

	// 第二返回值标识key是否存在
	val, flag := m1[4]
	fmt.Printf("map res: val=%v, flag=%v \n", val, flag)

	fmt.Printf("-------------\n")
	//3. 删除
	fmt.Printf("delete before m1=%v\n", m1)
	delete(m1, 2)
	val2, f := m1[2]
	fmt.Printf("map res: val2=%v, f=%v \n", val2, f)
	fmt.Printf("delete after m1=%v\n", m1)

}

func main() {

	test1()
}
