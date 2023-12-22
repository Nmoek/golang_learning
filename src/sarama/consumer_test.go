// Package sarama_test
// @Description: kafka客户端测试-消费者API1
package sarama_test

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

// @func: TestConsumer
// @date: 2023-12-17 04:14:42
// @brief: 消费API测试练习
// @author: Kewin Li
// @param t
func TestConsumer(t *testing.T) {

	cfg := sarama.NewConfig()

	group, err := sarama.NewConsumerGroup([]string{"localhost:9094"}, "test_group", cfg)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel() //cancel被调用就会让消费者停止消费

	start := time.Now()
	err = group.Consume(ctx, []string{"test_topic"}, &ConsumerHandler{})
	assert.NoError(t, err)

	t.Log("time cost=", time.Since(start).Milliseconds(), " ms")
}

type ConsumerHandler struct {
}

func (c *ConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("this is Setup \n")
	partitions := session.Claims()["test_topic"]
	offset := int64(0) // 从0开始执行

	//offset = sarama.OffsetNewest // 从broker最新消息开始
	//offset = sarama.OffsetOldest // 从broker最旧消息开始

	for _, part := range partitions {

		session.ResetOffset("test_topic", part, offset, "")
	}

	return nil
}

func (c *ConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Printf("this is Cleanup \n")
	return nil

}

// @func: ConsumeClaimV1
// @date: 2023-12-17 04:20:58
// @brief: 同步单个消费
// @author: Kewin Li
// @receiver c
// @param session
// @param claim
// @return error
func (c *ConsumerHandler) ConsumeClaimV1(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()

	for msg := range msgs {

		fmt.Printf("%v %v\n", string(msg.Key), string(msg.Value))
		// 提交
		session.MarkMessage(msg, "")
	}

	return nil
}

// @func: ConsumeClaim
// @date: 2023-12-17 04:21:11
// @brief: 异步批量消费
// @author: Kewin Li
// @receiver c
// @param session
// @param claim
// @return error
func (c *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	count := 0
	const batchSize = 10

	for {
		batch := make([]*sarama.ConsumerMessage, 0, 10)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		timeoutFlag := false

		// 并发处理 来一个处理一个
		var eg errgroup.Group
		for i := 0; i < batchSize && !timeoutFlag; i++ {

			select {
			case <-ctx.Done():
				// 已经超时
				fmt.Printf("time out! \n")
				timeoutFlag = true
			case msg, ok := <-msgs:
				// TODO: 考虑并发量低的情况，凑不够指定的阈值批量提交，需要设置超时
				if !ok {
					// channel已经被关闭
					cancel()
					fmt.Printf("channel closed \n")
					return nil
				}
				batch = append(batch, msg)

				//并发处理 来一个处理一个
				eg.Go(func() error {

					fmt.Printf("%v %v \n", string(msg.Key), string(msg.Value))

					return nil
				})
			}

			if timeoutFlag {
				break
			}
		}

		//并发处理 来一个处理一个
		err := eg.Wait()
		if err != nil {
			fmt.Printf("%v \n", err)
			continue
		}

		// 模拟业务批量处理
		// ...

		// 关闭channel
		cancel()
		// 批量提交
		for _, msg := range batch {
			session.MarkMessage(msg, "")
		}

		fmt.Printf("%d批%d个已经提交", count, len(batch))
		count++
	}

	return nil
}
