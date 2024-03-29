package main

import (
	"fmt"
	"os"
	"sort"
)

// @func: test1
// @brief: map的基本用法:创建、赋值、遍历、删除
// @author: Kewin Li
func test1() {

	//1. 创建、赋值
	m1 := map[int]string{}
	// m1 := make(map[int]string, 1)
	m1[3] = "333"
	m1[1] = "111"
	m1[2] = "222"

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

// @func: test2
// @date: 2023年8月23日
// @brief: map的键值排序问题
// @author: Kewin Li
func test2() {

	mp1 := map[int]string{}

	mp1[2] = "222"
	mp1[4] = "444"
	mp1[1] = "111"
	mp1[3] = "333"

	fmt.Printf("mp1: \n")
	fmt.Printf("mp1: %+v \n", mp1)

	keys := []int{}
	for m_k, _ := range mp1 {
		fmt.Printf("key=%v \n", m_k)
		keys = append(keys, m_k)
	}

	fmt.Printf("-------------\n")

	sort.Ints(keys)

	fmt.Printf("%+v \n", mp1[])
	for m_k, _ := range mp1 {
		fmt.Printf("key=%v \n", m_k)
	}
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("intput xxx [number]! \n")
		return
	}

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	}

}
