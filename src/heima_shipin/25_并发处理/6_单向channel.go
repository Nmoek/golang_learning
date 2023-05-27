package main

import (
	"fmt"
)

func producer(in chan<- int) {

	for i := 0; i <= 10; i++ {
		in <- i
	}

	close(in)
}

func comsumer(out <-chan int) {
	for num := range out {
		fmt.Printf("num=%d \n", num)
	}
}

func main() {

	m_ch := make(chan int, 0)

	var wrCh chan<- int = m_ch
	var rdCh <-chan int = m_ch

	go func() {
		num := <-rdCh
		fmt.Printf("test num=%d \n", num)
	}()

	wrCh <- 66666

	fmt.Printf("-------------------\n")

	go comsumer(m_ch)

	producer(m_ch)

}
