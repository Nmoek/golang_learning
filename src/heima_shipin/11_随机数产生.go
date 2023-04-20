package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(666)

	for i := 0; i < 5; i++ {
		fmt.Printf("rand= %d \n", rand.Int())
	}

}
