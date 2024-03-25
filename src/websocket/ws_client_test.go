package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8888/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		for {
			_, msg, err2 := conn.ReadMessage()
			if err2 != nil {
				fmt.Printf("读取服务器消息错误, %v \n", err2)
				return
			}
			fmt.Printf("收到来自服务器的消息, %s \n", string(msg))
		}
	}()

	ticker := time.NewTicker(time.Second)
	for now := range ticker.C {
		err2 := conn.WriteMessage(websocket.TextMessage, []byte(now.String()))
		if err2 != nil {
			fmt.Printf("向服务器发送消息出错, %v \n", err2)
		}
	}

}
