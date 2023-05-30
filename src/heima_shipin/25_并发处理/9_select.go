package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// @func: send_num
// @brief: 发送数据的协程
// @author: Kewin Li
// @param: chan int num_ch
// @param: chan bool quit_ch
func send_num(num_ch chan int, quit_ch chan bool, count int) {

	for i := 1; i < count; i++ {
		num := <-num_ch

		fmt.Printf("%d ", num)
	}

	fmt.Printf("\n")
	quit_ch <- true

}

// @func: test1
// @brief: 使用select打印斐波那契额数列
// @author: Kewin Li
func test1() {

	x := 1
	y := 1

	num_ch := make(chan int)
	quit_ch := make(chan bool)

	go send_num(num_ch, quit_ch, 10)

	for {
		select {
		case num_ch <- x:
			x, y = y, x+y
		case f := <-quit_ch:
			if f {
				fmt.Printf("send num finish!! \n")
				return
			}
		}
	}

}

// myHttpRes @brief: http请求的结果
type myHttpRes struct {
	url    string
	intval time.Duration
}

// @func: m_ping
// @brief: 自定义发起Http请求
// @author: Kewin Li
// @param: string url
// @return chan
func m_ping(url string) chan myHttpRes {

	m_ch := make(chan myHttpRes)

	t := time.Now()
	http.Get(url)
	m_ch <- myHttpRes{url, time.Since(t)}
	return m_ch
}

// @func: test2
// @brief: 用select比较http请求速度
// @author: Kewin Li
func test2() {

	url1 := "http://www.baidu.com/"
	url2 := "http://www.bilibili.com/"

	select {
	case res := <-m_ping(url1):
		fmt.Printf("http over %s, cost time=%d ms", res.url, res.intval.Milliseconds())
	case res := <-m_ping(url2):
		fmt.Printf("http over %s, cost time=%d ms", res.url, res.intval.Milliseconds())
	}

}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	}

}
