package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_NUMER_LEN = 1000

var count = 0

func m_swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// @func: popSort_v2
// @brief: 冒泡排序 优化版O(N^2)
// @author: Kewin Li
// @param: [MAX_NUMER_LEN]int numers
// @return [MAX_NUMER_LEN]int
func popSort_v2(numers [MAX_NUMER_LEN]int) [MAX_NUMER_LEN]int {

	f := true
	for i := 0; i < MAX_NUMER_LEN-1 && f; i++ {
		f = false
		for j := 0; j < MAX_NUMER_LEN-i-1; j++ {
			if numers[j] > numers[j+1] {
				count++
				m_swap(&numers[j], &numers[j+1])
				f = true
			}
		}

	}

	return numers
}

// @func: popSort_v1
// @brief: 冒泡排序 O(N^2)
// @author: Kewin Li
// @param: [10]int numer
// @return [10]int

func popSort_v1(numers [MAX_NUMER_LEN]int) [MAX_NUMER_LEN]int {

	for i := 0; i < MAX_NUMER_LEN-1; i++ {
		for j := 0; j < MAX_NUMER_LEN-1-i; j++ {
			if numers[j] > numers[j+1] {
				count++
				m_swap(&numers[j], &numers[j+1])
			}
		}
	}

	return numers
}

// @func: popSort_v0
// @brief: 冒泡排序- 错误版
// @author: Kewin Li
// @param: [MAX_NUMER_LEN]int numers
// @return [MAX_NUMER_LEN]int
func popSort_v0(numers [MAX_NUMER_LEN]int) [MAX_NUMER_LEN]int {

	for i := 0; i < MAX_NUMER_LEN; i++ {
		for j := i; j < MAX_NUMER_LEN; j++ {
			if numers[j] > numers[i] {
				count++
				m_swap(&numers[i], &numers[j])
			}
		}
	}

	return numers
}

func main() {

	numers := [MAX_NUMER_LEN]int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < MAX_NUMER_LEN; i++ {
		numers[i] = rand.Intn(1000)
	}

	// fmt.Printf("before numers=%v \n", numers)
	t1 := time.Now()
	numers = popSort_v1(numers)
	// fmt.Printf("after numers=%v \n", numers)

	fmt.Printf("count=%d \n", count)
	fmt.Printf("time=%v us\n", time.Since(t1).Microseconds())
}
