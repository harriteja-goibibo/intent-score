package wrapper

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goibibo/intent-score/internal/infrastructure/internal"
	"github.com/pkg/errors"
)

/*
	The whole point of kafka wrappers is that any sort of dealing with actual kafka package i.e types.variables
	etc which should be independent of application logic.
	What we gain ? -> In future when if we plan to shift back to sarama instead of confluent, you just change the
	code in wrapper just the way the newer package expects, without touching a single line of application code.
*/
func NewConsumerClient(consumerImpl internal.KafkaConsumer) (*ConsumerClient, error) {
	if consumerImpl == nil {
		return nil, errors.New("Consumer Implementation for kafka is nil")
	}
	return &ConsumerClient{
		Kafka: consumerImpl,
	}, nil
}

func (c *ConsumerClient) Poll(timeout int, successEvent chan<- KafkaEventResponse, failureEvent chan<- error) {
	event := c.Kafka.Poll(timeout)
	switch e := event.(type) {
	case *kafka.Message:
		successEvent <- KafkaEventResponse{
			RawData:   e.Value,
			Topic:     *e.TopicPartition.Topic,
			Partition: e.TopicPartition.Partition,
			Offset:    int32(e.TopicPartition.Offset),
		}
	case kafka.Error:
		failureEvent <- errors.New(e.String())
	}
}
func (c *ConsumerClient) SubscribeTopics(topics []string, rebalanceCBRequired bool) error {
	if rebalanceCBRequired {
		// TODO;
	}
	return c.Kafka.SubscribeTopics(topics, nil)
}

func (c *ConsumerClient) Commit() (response []KafkaEventResponse, err error) {
	commitResponse, commitErr := c.Kafka.Commit()
	if err != nil {
		return nil, commitErr
	}
	for _, topicResponse := range commitResponse {
		response = append(response, KafkaEventResponse{
			Topic:     *topicResponse.Topic,
			Partition: topicResponse.Partition,
			Offset:    int32(topicResponse.Offset),
		})
	}
	return response, nil
}

// KafkaConsumer creates the actual kafka consumer of type confluent client, which might be changed in future.
func KafkaConsumer(config ConsumerConfig) (internal.KafkaConsumer, error) {
	confluentConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"heartbeat.interval.ms":              config.HeartBeatInterval, // For consumer to specify its alive to the broker
		"session.timeout.ms":                 config.SessionTimeout,    // If the broker doesn't receive heartbeat request from consumer within this timeout, it will delete the consumer and trigger rebalance
		"retries":                            config.MaxRetries,
		"retry.backoff.ms":                   config.IntervalBetweenRetries,  // If a request is failed, this will be the timeout before kafka attempt to retry
		"topic.metadata.refresh.interval.ms": config.MetadataRefreshInterval, // At this interval, topic and broker data are refreshed. (Any new leaders, new brokers were picked up)
		"bootstrap.servers":                  config.BootstrapServers,
		"group.id":                           config.ConsumerGroupID,
		"partition.assignment.strategy":      config.PartitionStrategy,
		"enable.auto.commit":                 config.AutoCommit,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Error creating kafka consumer via confluent client")
	}
	return confluentConsumer, nil
}
