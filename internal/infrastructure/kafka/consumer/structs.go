package consumer

import (
	"github.com/goibibo/intent-score/internal/infrastructure/kafka/wrapper"
)

type KafkaConsumerWrapper interface {
	Poll(timeout int, successEvent chan<- wrapper.KafkaEventResponse, failureEvent chan<- error)
	SubscribeTopics(topics []string, rebalanceCBRequired bool) error
	Commit() ([]wrapper.KafkaEventResponse, error)
}
type Client struct {
	KafkaConsumer KafkaConsumerWrapper
}
