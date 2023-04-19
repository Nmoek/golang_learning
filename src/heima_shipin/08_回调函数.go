package main

import (
	"fmt"
)

type FuncType func(int, int) int

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

// Go中最简单的多态
func Cal(a int, b int, f FuncType) (res int) {
	return f(a, b)
}

func main() {

	fmt.Printf("add res=%d \n", Cal(1, 2, Add))
	fmt.Printf("sub res=%d \n", Cal(5, 3, Sub))
	fmt.Printf("mul res=%d \n", Cal(2, 3, Mul))
}
