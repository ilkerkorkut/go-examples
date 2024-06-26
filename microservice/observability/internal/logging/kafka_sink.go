package logging

import (
	"time"

	"github.com/IBM/sarama"
)

type kafkaSink struct {
	producer sarama.SyncProducer
	topic    string
}

func (s kafkaSink) Write(b []byte) (int, error) {
	_, _, err := s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: s.topic,
		Key:   sarama.StringEncoder(time.Now().String()),
		Value: sarama.ByteEncoder(b),
	})
	return len(b), err
}

func (s kafkaSink) Sync() error {
	return nil
}

func (s kafkaSink) Close() error {
	return nil
}
