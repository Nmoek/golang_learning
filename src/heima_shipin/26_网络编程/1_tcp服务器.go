package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	// 1. 监听
	listener, err1 := net.Listen("tcp", "192.168.77.136:8888")
	if err1 != nil {
		fmt.Printf("start listen err! %s", err1.Error())
		return
	}

	fmt.Printf("start listen...\n")

	for {
		// 2. 获取连接
		new_con, err2 := listener.Accept()
		if err2 != nil {
			fmt.Printf("accept connect err! %s", err2.Error())
			return
		}

		fmt.Printf("get new connect ip/port[%s] \n", new_con.RemoteAddr().String())

		go func(con net.Conn) {
			ip := con.RemoteAddr().String()
			for {

				buf := make([]byte, 100)
				n, err := con.Read(buf)
				if err != nil && err != io.EOF {
					fmt.Printf("[%s] read err! %s \n", ip, err.Error())
					continue
				}

				if n == 0 {
					fmt.Printf("[%s] close! \n", ip)
					return
				}

				fmt.Printf("[%s] read n=%d, buf:\n %s\n", ip, n, buf)

				if string(buf) == "quit" {
					fmt.Printf("[%s] quit \n", ip)
					return
				}
			}

		}(new_con)
	}
}
