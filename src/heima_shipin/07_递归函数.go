package main

import (
	"fmt"
	"time"
)

func m_func1(val int) int {

	if val == 0 {
		return 0
	}

	fmt.Printf("val=%d \n", val)

	return m_func1(val - 1)

}

func m_func2(val int) int {

	if val == 0 {
		return 0
	}

	return val + m_func2(val-1)
}

func m_func3(val int) (sum int) {

	for i := 0; i <= val; i++ {
		sum += i
	}

	return sum
}

func main() {

	m_func1(5)
	fmt.Printf("----------------------\n")
	t1 := time.Now()
	fmt.Printf("sum=%d \n", m_func2(1000000))
	fmt.Printf("t1=%v \n", time.Since(t1).Milliseconds()) //373ms
	fmt.Printf("----------------------\n")

	t2 := time.Now()
	fmt.Printf("sum=%d \n", m_func3(1000000))
	fmt.Printf("t1=%v \n", time.Since(t2).Milliseconds()) // <1ms

	//递归性能消耗挺严重

}
