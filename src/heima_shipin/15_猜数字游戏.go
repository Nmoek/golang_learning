package main

import (
	"fmt"
	"math/rand"
	"time"
)

var m_num int = 0 //拿到一个四位数

func guessNum(num int, tar_num int) (res bool) {

	var guessInfo string = ""
	res = true

	for i := 0; i < 4; i++ {
		var tmp string = ""

		if num%10 > tar_num%10 {
			res = false
			tmp = fmt.Sprintf("第%d位数字大了\n", 4-i)
		} else if num%10 < tar_num%10 {
			res = false
			tmp = fmt.Sprintf("第%d位数字小了\n", 4-i)
		} else {
			tmp = fmt.Sprintf("第%d位数字正确\n", 4-i)
		}

		guessInfo += tmp
		num /= 10
		tar_num /= 10
	}

	if !res {
		fmt.Printf("%v", guessInfo)
	} else {
		fmt.Printf("恭喜猜中了!\n")
	}

	return res
}

func main() {

	var test_num int

	rand.Seed(time.Now().UnixNano())

	for {
		m_num = rand.Intn(10000)
		if m_num >= 1000 {
			break
		}
	}

	for {
		fmt.Printf("input u num: ")
		fmt.Scanf("%d", &test_num)
		if test_num < 1000 || test_num >= 10000 {
			fmt.Printf("请输入四位数!\n")
			continue
		}

		if guessNum(test_num, m_num) {
			break
		}

		fmt.Printf("------------------- \n")
	}

}
