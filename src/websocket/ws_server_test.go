package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	upgrader := websocket.Upgrader{}

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			writer.Write([]byte("初始化websocket失败"))
			return
		}

		ws := Ws{conn: conn}
		go func() {
			ws.ReadCycle()
		}()

		go func() {
			ticker := time.NewTicker(time.Second)
			for now := range ticker.C {
				ws.Write(now.String())
			}
		}()
	})

	http.ListenAndServe(":8888", nil)
}

type Ws struct {
	conn *websocket.Conn
}

func (w *Ws) ReadCycle() {
	for {
		msgType, msg, err := w.conn.ReadMessage()
		if err != nil {
			fmt.Printf("读取客户端消息出错, %v \n", err)
			return
		}
		switch msgType {
		case websocket.CloseMessage:
			fmt.Printf("[%s]客户端关闭 \n", w.conn.RemoteAddr().String())
			w.conn.Close()
		case websocket.TextMessage, websocket.BinaryMessage:
			fmt.Printf("收到客户端消息, %s \n", string(msg))
		default:

		}
	}
}

func (w *Ws) Write(msg string) {
	err := w.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("向客户端发送消息失败, %v \n", err)
	}
}
