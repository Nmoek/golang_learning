package main

import (
	"fmt"
	"time"
)

// 注意在这个例子中，公共资源应该是：标准输出文件句柄（即屏幕显示）

var m_chan = make(chan int)

func m_printf(data string) {

	for _, d := range data {
		fmt.Printf("%c", d)
		time.Sleep(30 * time.Millisecond)
	}
}

func m_go1() {

	m_printf("hello")
	m_chan <- 666
}

func m_go2() {

	<-m_chan //管道中没有数据时会阻塞
	m_printf("world!")
	m_chan <- 666

}

func m_go3() {
	<-m_chan
	m_printf("ljk6666!")

}

func main() {

	go m_go1()
	go m_go2()
	go m_go3()

	for {

	}

}
