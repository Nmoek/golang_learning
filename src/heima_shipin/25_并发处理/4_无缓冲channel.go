package main

import (
	"fmt"
	"time"
)

func main() {

	m_chan := make(chan int, 0)
	fmt.Printf("len=%d, cap=%d \n", len(m_chan), cap(m_chan))

	go func() {

		for i := 0; i < 2; i++ {
			fmt.Printf("im sub goroutine i=%d \n", i)
			m_chan <- i
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 2; i++ {
		num := <-m_chan
		fmt.Printf("im main goroutine, num=%d \n", num)
	}
}
