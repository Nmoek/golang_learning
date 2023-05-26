package main

import (
	"fmt"
	"time"
)

// 注意在这个例子中，公共资源应该是：标准输出文件句柄（即屏幕显示）

func m_printf(data string) {

	for _, d := range data {
		fmt.Printf("%c", d)
		time.Sleep(30 * time.Millisecond)
	}
}

func m_go1() {

	m_printf("hello")

}

func m_go2() {

	m_printf("world!")

}

func main() {

	go m_go1()
	go m_go2()

	for {

	}

}
