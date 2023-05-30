package main

import (
	"fmt"
	"net"
	"time"
)

var MY_MAX_CON_NUM int = 10

func main() {

	cons := make([]net.Conn, 0, MY_MAX_CON_NUM)
	notify_ch := make(chan bool)
	quit_count := 0

	// 1. 连接
	for i := 0; i < MY_MAX_CON_NUM; i++ {
		con, err1 := net.Dial("tcp", "192.168.77.136:8888")
		if err1 != nil {
			fmt.Printf("[%d] net Dial err! %s\n", i, err1.Error())
			continue
		}
		cons = append(cons, con)

		fmt.Printf("[%d]connect server sccuss! local[%s], remote[%s]\n", i, con.LocalAddr().String(), con.RemoteAddr().String())

	}

	// 2. 发送内容
	for i := 0; i < len(cons); i++ {
		go func(con net.Conn, idx int) {
			var buf string
			buf = fmt.Sprintf("hello, im %d", idx)
			for j := 0; j < 50; j++ {
				_, err := con.Write([]byte(buf))
				if err != nil {
					fmt.Printf("[%s] write err! %s \n", con.LocalAddr().String(), err.Error())
					continue
				}

				time.Sleep(100 * time.Millisecond)
			}

			con.Write([]byte("quit"))
			con.Close()

			notify_ch <- true

		}(cons[i], i)
	}

	for {
		f := <-notify_ch
		if f {
			quit_count++
			if quit_count >= len(cons) {
				break
			}
		}
	}

}
