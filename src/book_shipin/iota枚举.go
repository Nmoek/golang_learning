package main

import (
	"fmt"
)

func main() {
	
	// 1. 标准用法
	// 每隔一行会自动累加1
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a= %d, b= %d, c=%d\n", a, b, c)

	// 2. 遇到const就重置为0
	const (
		a1 = iota
		b1 = iota
	)
    // 错误示范 只能在const域中使用
	// c1 := iota
	const d1 = iota 

	fmt.Printf("a1=%d, b1=%d, d1=%d \n", a1, b1, d1)

}