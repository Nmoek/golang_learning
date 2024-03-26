// Package saramax
// @Description: 同步消费-批量提交
package saramax

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"strconv"
	"time"
)

// 批量提交阈值
const batchSize = 10

type BatchHandler[T any] struct {
	fn func(msgs []*sarama.ConsumerMessage, ts []T) error
}

func NewBatchHandler[T any](fn func(msgs []*sarama.ConsumerMessage, ts []T) error) *BatchHandler[T] {
	return &BatchHandler[T]{
		fn: fn,
	}
}

func (b *BatchHandler[T]) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BatchHandler[T]) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (b *BatchHandler[T]) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	msgs := claim.Messages()

	for {
		batch := make([]*sarama.ConsumerMessage, 0, batchSize)
		ts := make([]T, 0, batchSize)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		timeoutFlag := false
		for i := 0; i < batchSize && timeoutFlag; {
			select {
			case <-ctx.Done():
				//TODO: 会话超时, 日志埋点
				timeoutFlag = true
			case msg, ok := <-msgs:
				// channel被关闭
				if !ok {
					cancel()
					// TODO: 通道关闭错误
					return nil
				}
				var t T
				err := json.Unmarshal(msg.Value, &t)
				if err != nil {
					fmt.Println("反序列化消息失败", err, "topic="+msg.Topic,
						"partition="+strconv.FormatInt(int64(msg.Partition), 10), "offset="+strconv.FormatInt(msg.Offset, 10))

					continue
				}

				batch = append(batch, msg)
				ts = append(ts, t)
				// 严格不漏记这里进行递增
				i++
			}

		}

		cancel()

		err := b.fn(batch, ts)
		if err != nil {

			fmt.Println("消息业务处理出错", err, ts)

		}

		for _, msg := range batch {
			session.MarkMessage(msg, "")
		}
	}

}
