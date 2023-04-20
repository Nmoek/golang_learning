package main

import (
	"Cal"
	"fmt"
	t1 "test1"
)

func main() {
	t1.Test1_1()
	t1.Test1_2()

	res := Cal.Add(1, 2)
	fmt.Printf("add res=%d \n", res)
}
