package wrapper

import "github.com/goibibo/intent-score/internal/infrastructure/internal"

type ConsumerConfig struct {
	HeartBeatInterval       int
	SessionTimeout          int
	MaxRetries              int
	IntervalBetweenRetries  int
	MetadataRefreshInterval int
	BootstrapServers        string
	ConsumerGroupID         string
	PartitionStrategy       string
	AutoCommit              bool
}
type KafkaEventResponse struct {
	RawData   []byte
	Topic     string
	Partition int32
	Offset    int32
}

type ConsumerWrap interface {
	Poll(timeout int, successEvent chan<- KafkaEventResponse, failureEvent chan<- error)
	SubscribeTopics(topics []string, rebalanceCBRequired bool) error
	Commit() (response []KafkaEventResponse, err error)
}

type ConsumerClient struct {
	Kafka internal.KafkaConsumer
}
