package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {

		for i := 0; i < 5; i++ {
			fmt.Printf("im sub goroutine! \n")
			// 让当前协程直接退出
			runtime.Goexit()
		}

	}()

	// 让主goroutine让出
	runtime.Gosched()

	for i := 0; i < 2; i++ {
		fmt.Printf("im main goroutine! \n")
	}

}
