package simpleim

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/ecodeclub/ekit/syncx"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"ws_test/simpleim/saramax"

	"time"
)

type WsGateway struct {
	svc        *Service
	client     sarama.Client
	conns      *syncx.Map[int64, *Conn]
	instanceId string
	upgrader   *websocket.Upgrader
}

func NewWsGateway(svc *Service, client sarama.Client, instanceId string) *WsGateway {
	return &WsGateway{
		svc:        svc,
		client:     client,
		conns:      &syncx.Map[int64, *Conn]{},
		instanceId: instanceId,
		upgrader:   &websocket.Upgrader{},
	}
}

func (w *WsGateway) Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", w.wsHandler)

	err := w.subscribeMsg()
	if err != nil {
		return err
	}

	return http.ListenAndServe(addr, mux)
}

func (w *WsGateway) wsHandler(writer http.ResponseWriter, request *http.Request) {

	c, err := w.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		writer.Write([]byte("初始化websocket失败"))
		return
	}

	// 拿到uid
	// 从jwt的token拿到或者session中获取
	uid := w.mockUid(request)
	conn := &Conn{c}

	w.conns.Store(uid, conn)
	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				return
			}

			//转发后端
			var msg Message
			err = json.Unmarshal(data, &msg)
			if err != nil {
				continue
			}

			// 业务转发可能会阻塞, 异步处理
			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				err = w.svc.Receive(ctx, uid, msg)
				cancel()
				if err != nil {
					err = conn.Send(Message{
						Type:    "result",
						Content: "FAILED",
						Seq:     msg.Seq,
					})
					if err != nil {
						//TODO: 记录日志
					}
				}
			}()

		}
	}()
}

func (w *WsGateway) subscribeMsg() error {
	group, err := sarama.NewConsumerGroupFromClient(w.instanceId, w.client)
	if err != nil {
		return err
	}

	go func() {
		err2 := group.Consume(context.Background(), []string{eventName}, saramax.NewHandler[Event](w.Consume))
		if err2 != nil {

		}
	}()

	return nil
}

// @func: Consume
// @date: 2024-03-26 22:48:37
// @brief: 消费者
// @author: Kewin Li
// @receiver w
// @param msg
// @param event
// @return error
func (w *WsGateway) Consume(msg *sarama.ConsumerMessage, event Event) error {
	conn, ok := w.conns.Load(event.Receiver)
	if !ok {
		return nil
	}

	fmt.Printf("收到消息: %v \n", event)

	// 消息转发
	return conn.Send(event.Msg)
}

// @func: mockUid
// @date: 2024-03-26 22:29:31
// @brief: 模拟拿到用户ID
// @author: Kewin Li
// @receiver w
// @param req
// @return int64
func (w *WsGateway) mockUid(req *http.Request) int64 {
	uid, _ := strconv.ParseInt(req.Header.Get("uid"), 10, 64)
	return uid
}

type Conn struct {
	*websocket.Conn
}

func (c *Conn) Send(msg Message) error {
	val, _ := json.Marshal(msg)

	return c.WriteMessage(websocket.TextMessage, val)
}

// 前后端交互数据协议
type Message struct {
	Seq string `json:"seq,omitempty"`
	// {"type": "image", content:"http://myimage"}
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	// channel ID
	Cid int64 `json:"cid,omitempty"`
}
