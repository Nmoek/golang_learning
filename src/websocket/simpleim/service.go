package simpleim

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"strconv"
)

type Service struct {
	producer sarama.SyncProducer
}

func (s *Service) Receive(ctx context.Context, sender int64, msg Message) error {
	members := s.mockFindMember()

	for _, m := range members {
		// 不需要给自己转发
		if sender == m {
			continue
		}
		event := Event{
			Msg:      msg,
			Receiver: m,
		}
		val, _ := json.Marshal(event)
		_, _, err := s.producer.SendMessage(&sarama.ProducerMessage{
			Topic: eventName,
			Key:   sarama.StringEncoder(strconv.FormatInt(m, 10)),
			Value: sarama.StringEncoder(val),
		})
		if err != nil {
			continue
		}

	}
	return nil
}

// @func: mockFindMember
// @date: 2024-03-26 22:21:40
// @brief: 模拟发现群姐成语
// @author: Kewin Li
// @receiver s
func (s *Service) mockFindMember() []int64 {
	return []int64{1, 2, 3, 4}
}
