package main

import (
	"fmt"
	"os"
	"time"
)

// @func: test1
// @brief: Timer只会触发一次
// @author: Kewin Li
func test1() {

	timer := time.NewTimer(2 * time.Second)

	fmt.Printf("cur time=%v \n", time.Now())

	res := <-timer.C
	fmt.Printf("res=%v \n", res)
}

func task1() {

	fmt.Printf("hello world!\n")

}

// @func: test2
// @brief: Timer 添加异步回调
// @author: Kewin Li
func test2() {

	res := time.AfterFunc(2*time.Second, task1)
	fmt.Printf("res=%T %v \n", res, *res)
}

// @func: test3
// @brief: Timer定时器停止与重置
// @author: Kewin Li
func test3() {

	timer := time.NewTimer(3 * time.Second)

	go func() {
		res := <-timer.C

		fmt.Printf("im sub gorotuine! res=%v \n", res)

		os.Exit(0)
	}()

	timer.Stop()
	timer.Reset(1 * time.Second)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	case '3':
		test3()
	}

	for {

	}

}
