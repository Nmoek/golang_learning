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

func main() {

	func1(1)

	fmt.Printf("---------------------- \n")

	func1(1, 2, 3)

	fmt.Printf("---------------------- \n")

}
