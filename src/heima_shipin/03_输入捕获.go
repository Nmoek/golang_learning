package main

import (
	"fmt"
)

func main() {

	var a int
	fmt.Scanf("%d", &a)
	//input: 121231
	fmt.Printf("a=%d \n", a)
	// output:a=121231

	var b int
	fmt.Scan(&b)
	//input: 1.3651651
	fmt.Printf("b=%f \n", b)
	// output:b=%!f(int=1) 注意：只取了一个整数 1

	var c float64
	fmt.Scanf("%f", &c)
	//注意: 此时输入缓冲区中还有内容 .3651651
	fmt.Printf("c=%f \n", c)
	// output:c=3651651.000000

}
