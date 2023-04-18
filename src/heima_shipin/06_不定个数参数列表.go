package main

import (
	"fmt"
)

func func1(args ...int) {
	fmt.Printf("aegs len=%d \n", len(args))

	for _, val := range args {
		fmt.Printf("%d ", val)
	}

	fmt.Printf("\n")
}

func func2_do(tmp ...int) {

	fmt.Printf("tmp len=%d \n", len(tmp))

	for _, val := range tmp {

		fmt.Printf("%d ", val)
	}

	fmt.Printf("\n")

}

func func2(args ...int) {

	func2_do(args...) //将本次不定参数列表全部传递

	func2_do(args[:2]...) //切片形式部分传递 第一个元素~1

	func2_do(args[2:]...) //切片形式部分传递 2~最后一个元素
}

func main() {

	func1(1)

	fmt.Printf("---------------------- \n")

	func1(1, 2, 3)

	fmt.Printf("---------------------- \n")

	func2(1, 2, 3, 4, 5, 6)

}
