package main

import (
	"fmt"
)

func main() {

	// 基本用法：
	// for 初始条件 ; 判断条件 ; 条件变化 {
	//}
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("sum=%d \n", sum)

	fmt.Printf("--------------------- \n")

	// range迭代用法:
	// rang默认返回两个值: 1. 从0开始的元素位置 2. 元素本身
	str := "abcdef"
	for idx, val := range str {
		fmt.Printf("idx=%d, val=%c \n", idx, val)
	}

}
