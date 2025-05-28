package kk

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	reader *kafka.Reader
}

func (r *KafkaReader) ReadMsg(ctx context.Context) (kafka.Message, error) {
	return r.reader.ReadMessage(ctx)
}

func (r *KafkaReader) Close() {
	if r.reader != nil {
		r.reader.Close()
	}
}

func GetReader(brokers []string, groupId, topic string) *KafkaReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupId, //同一个组下的consumer 协同工作 共同消费topic队列中的内容
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	k := &KafkaReader{reader: reader}
	return k
}
