package main

import (
	"fmt"
	"os"
	"time"
)

func test1() {

	timer := time.NewTicker(3 * time.Second)

	go func() {

		for {
			res := <-timer.C

			fmt.Printf("hello11111 %v\n", res)
		}

	}()

	go func() {

		for {
			res := <-timer.C

			fmt.Printf("hello222222 %v\n", res)
		}

	}()

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()

	}

	for {

	}

}
