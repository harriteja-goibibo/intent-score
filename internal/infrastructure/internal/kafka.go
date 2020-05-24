package internal

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer interface {
	Poll(timeout int) kafka.Event
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error
	Commit() ([]kafka.TopicPartition, error)
}
