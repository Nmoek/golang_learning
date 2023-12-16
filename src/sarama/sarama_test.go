// Package sarama_test
// @Description: kafka客户端测试
package sarama_test

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var addr = []string{"localhost:9094"}

// @func: TestSyncProducer
// @date: 2023-12-16 23:03:17
// @brief: 同步发送接口
// @author: Kewin Li
// @param t
func TestSyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addr, cfg)
	assert.NoError(t, err)

	cf

	/*向哪个分区发送消息*/
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner // 轮询
	//cfg.Producer.Partitioner = sarama.NewRandomPartitioner            // 随机
	//cfg.Producer.Partitioner = sarama.NewHashPartitioner              // 根据key的hash值进行选择
	//cfg.Producer.Partitioner = sarama.NewManualPartitioner            // 根据Message中指定的分区
	//cfg.Producer.Partitioner = sarama.NewConsistentCRCHashPartitioner // 一致性hash 使用CRC16算法
	//cfg.Producer.Partitioner = sarama.NewCustomPartitioner()          // 自定义hash

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test_topic",
		Key:   sarama.StringEncoder("这是测试同步消息的key"),
		Value: sarama.StringEncoder("这是一条测试同步消息"),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("key1"),
				Value: []byte("val1"),
			},
		},
		Metadata: "metadata",
	})

	assert.NoError(t, err)
}

// @func: TestAsyncProducer
// @date: 2023-12-16 23:03:09
// @brief: 异步发送接口
// @author: Kewin Li
// @param t
func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(addr, cfg)
	assert.NoError(t, err)

	go func() {
		for {
			select {
			case msg := <-producer.Successes():
				t.Log("消息发布成功", string(msg.Value.(sarama.StringEncoder)))
			case err2 := <-producer.Errors():
				val, _ := err2.Msg.Value.Encode()
				t.Error("消息发送失败", err2.Err, "val=", string(val))
			}

			time.Sleep(time.Second)
		}

	}()

	msgs := producer.Input()
	count := int64(0)
	for {

		msgs <- &sarama.ProducerMessage{
			Topic: "test_topic",
			Key:   sarama.StringEncoder("这是测试异步消息的key"),
			Value: sarama.StringEncoder("这是一条测试异步消息" + strconv.FormatInt(count, 10)),
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("key2"),
					Value: []byte("val2"),
				},
			},
			Metadata: "metadata",
		}
		count++
		time.Sleep(time.Second)

	}

}
