package main

import (
	"fmt"
)

func main() {

	//分配5个int长度的空间
	slice := make([]int, 5)

	fmt.Printf("slice=%v  len=%d \n", slice, len(slice))

	//分配5个int长度的空间 预留10个int长度的空间
	slice = make([]int, 5, 10)

	fmt.Printf("slice=%v  len=%d \n", slice, len(slice))

	//error: 不能给变量分配空间
	// a := make(int, 1)

}
