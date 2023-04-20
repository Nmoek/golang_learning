package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 每次设置一个不同的时间种子
	// 注意: 伪随机，会出现重复的数
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		// 限定在100内的随机数
		fmt.Printf("rand= %d \n", rand.Intn(100))
	}

}
