package main

import (
	"./test_number"
	"fmt"
)

func main() {

	var cmd byte
	fmt.Scan("%c", &cmd)

	fmt.Printf("cmd=%c'n", cmd)
	switch cmd {
	/*测试整数*/
	case '0':
		{
			test_int()
			break
		}

	}

	fmt.Printf("\n!!! variable test over !!!\n")

}
