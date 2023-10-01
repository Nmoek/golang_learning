/*
 * @file:
 * @brief: 有关数组的练习题
 * @author: Kewin Li
 * @date: 2023年7月21日
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type m_data []int

func (data m_data) Len() int               { return len(data) }
func (data m_data) Swap(i int, j int)      { data[i], data[j] = data[j], data[i] }
func (data m_data) Less(i int, j int) bool { return data[i] < data[j] }

// @func: timu_find55
// @brief: 题目1-找出是否存在55
// @author: Kewin Li
// @param: [10]int nums
func timu_find55(nums [10]int) (bool, int) {

	sort.Sort(m_data(nums[:]))

	idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= 55 })
	if idx < len(nums) && nums[idx] == 55 {
		return true, idx
	} else {
		return false, -1
	}

}

// @func: timu1_Cal
// @brief: 题目1-求平均值、最大值+最小值的下标
// @author: Kewin Li
// @param: [10]int nums
func timu1_Cal(nums [10]int) {

	total := 0
	max_idx := -1
	max_val := -1
	min_idx := -1
	min_val := 200

	for i, v := range nums {
		total += v
		if v > max_val {
			max_val = v
			max_idx = i
		}

		if v < min_val {
			min_val = v
			min_idx = i
		}
	}

	fmt.Printf("av=%.2f, max_idx=%d, min_idx=%d \n", float64(total)/float64(10), max_idx, min_idx)

}

// @func: timu1_printWithReverse
// @brief: 题目1-倒序打印
// @author: Kewin Li
// @param: [10]int nums
func timu1_printWithReverse(nums [10]int) {

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("%d ", nums[i])
	}

	fmt.Printf("\n")
}

// @func: test1
// @brief:
// 随机生成10个整数(1~100)保存到数组，并倒序打印以及求平均值；求出最大值和最小值的下标，并查找其中是否包含55
// @author: Kewin Li
func timu1() {

	rand.Seed(time.Now().UnixMicro())

	var nums [10]int

	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(99) + 1
	}

	fmt.Printf("nums=%v \n", nums)
	fmt.Printf("----------------------\n")
	// 1. 倒序打印
	timu1_printWithReverse(nums)

	fmt.Printf("----------------------\n")
	//2. 求平均值、最大值+最小值的下标
	timu1_Cal(nums)

	fmt.Printf("----------------------\n")
	// 3. 找出是否存在55
	b, idx := timu_find55(nums)
	if b {
		fmt.Printf("55 find!! index=%d\n", idx)
	} else {
		fmt.Printf("55 not find! \n")
	}

}

func main() {

	args := os.Args

	switch args[1][0] {

	case '1':
		timu1()

	}

}
