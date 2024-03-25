package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"testing"
)

type Hub struct {
	lock  sync.Mutex
	conns map[string]*websocket.Conn
	sync.Map
}

func newHub() *Hub {
	return &Hub{
		conns: map[string]*websocket.Conn{},
		lock:  sync.Mutex{},
	}
}

func (h *Hub) addConn(key string, conn *websocket.Conn) {
	h.lock.Lock()
	h.conns[key] = conn
	h.lock.Unlock()

	go func() {
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("[%s]接收数据失败, %v \n\n", conn.RemoteAddr().String(), err)
			}

			h.lock.Lock()
			for n, c := range h.conns {
				if n == key {
					continue
				}

				err = c.WriteMessage(msgType, msg)
				if err != nil {
					fmt.Printf("向[%s]发送数据失败, %v\n", c.RemoteAddr().String(), err)
				}
			}

			h.lock.Unlock()
		}
	}()
}
func (h *Hub) getConn(key string) (conn *websocket.Conn, err error) {
	h.lock.Lock()
	c, ok := h.conns[key]
	if !ok {
		return nil, fmt.Errorf("没有该连接, %s \n", key)
	}
	h.lock.Unlock()

	return c, nil
}

func TestHub(t *testing.T) {
	h := newHub()
	upgrader := websocket.Upgrader{}
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			panic(err)
		}
		name := request.URL.Query().Get("name")

		h.addConn(name, conn)

	})

	http.ListenAndServe(":8888", nil)
}
