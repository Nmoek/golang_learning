package main

import (
	"fmt"
)

func test1(num int) {
	/*
	 * go的switch会默认加break
	 */
	switch num {
	case 1:
		{
			fmt.Printf("11111111 \n")
		}
		fallthrough
		//新关键字 fallthrough 无条件继续执行下一个case

	case 2:
		fmt.Printf("22222222 \n")
	}

}

func test2(num int) {

	//支持加入一个初始化语句
	switch n := 2; n {
	case 1:
		{
			fmt.Printf("1111111 \n")
		}
	default:
		fmt.Printf("this other option!! \n")
	}
}

func test3() {

	score := 85
	switch { //可以不写条件, 引用外部变量
	case score > 90:
		{
			fmt.Printf("优秀!\n")
		}
	case score >= 85:
		{
			fmt.Printf("良好!\n")
		}
	}
}

func main() {

	var num int
	fmt.Scan(&num)

	// test1(num)

	// test2(num)
	test3()

	fmt.Printf("over! \n")
}
