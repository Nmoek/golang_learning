package main

import (
	"fmt"
	"os"
	"regexp"
)

// TODO： 需要系统学习使用场景以及常用的正则表达式

// @func: test3
// @brief: 提取网页中的信息
// @author: Kewin Li
func test3() {

	buf := `
	<html>
	<div>哈哈
	6666
	sdfsdf
	</div>
	<div>6666</div>
	<div>20561581</div>
	</html>
	`

	reg := regexp.MustCompile(`<div>(?s:.*?)</div>`)

	res := reg.FindAllStringSubmatch(buf, -1)

	fmt.Printf("res=%+v \n", res)

}

// @func: test2
// @brief: 找出字符串中所有的小数
// @author: Kewin Li
func test2() {
	buf := "5416. 31.5196 .1 sdf.16sd ..sdf -3.15 7. 0.565"

	reg := regexp.MustCompile(`-?\d+\.\d+`)

	res := reg.FindAllStringSubmatch(buf, -1)

	fmt.Printf("res=%+v \n", res)

}

// @func: test1
// @brief: 正则包的基本使用
// @author: Kewin Li
func test1() {

	buf := "asc aic a5c 886 464 cia"

	// 1. 创建一个解释器
	reg := regexp.MustCompile(`a\dc`)

	// 2. 根据规则进行结果提取
	res := reg.FindAllStringSubmatch(buf, -1)

	fmt.Printf("res=%+v \n", res)

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	case '3':
		test3()
	}

}
