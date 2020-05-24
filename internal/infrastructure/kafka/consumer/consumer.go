package consumer

import (
	"errors"
	"fmt"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka/wrapper"
	"log"
)

// New creates the client that resolves dependencies and exposes functionality of this package
func New(kafkaConsumer KafkaConsumerWrapper) *Client {
	return &Client{
		KafkaConsumer: kafkaConsumer,
	}
}

func pushEventToJobsQueue(jobs chan<- kafka.Consumer, isEventProcessed chan kafka.Consumer,
	errorChan chan error) {
	go func() {
		for {
			select {
			case newJob := <-isEventProcessed:
				jobs <- newJob
			case err := <-errorChan:
				log.Println(err)
			}
		}
	}()
}

func unmarshalAndNotifyJobsQueue(consumer kafka.Consumer, triggerUnmarshal chan wrapper.KafkaEventResponse,
	isEventProcessed chan kafka.Consumer, errChan chan error) {
	defer func() {
		if r := recover(); r != nil {
			errStr := fmt.Sprintf("had a panic attack .. content format suspected .. %s", r)
			errChan <- errors.New(errStr)
		}
	}()
	for {
		select {
		case kafkaStreamData := <-triggerUnmarshal:
			if len(kafkaStreamData.RawData) > 0 {
				if data, err := consumer.Unmarshal(map[string][]byte{"data": kafkaStreamData.RawData, "topic": []byte(kafkaStreamData.Topic)}); err != nil {
					errChan <- errors.New("Error unmarshalling data -> " + err.Error())
				} else {
					isEventProcessed <- data
				}
			}
		}
	}
}

/*
	listen starts up a go routine which keeps polling kafka regularly and checks for an event.
	If there is an event, it pushed the raw bytes of data, the topic from which its consumed, partition etc.
	In case there is an error while consuming

*/

func listenAndTriggerUnmarshal(kafkaConsumer KafkaConsumerWrapper, responseChan chan wrapper.KafkaEventResponse, errorChan chan error) {
	successEvent := make(chan wrapper.KafkaEventResponse)
	errorEvent := make(chan error)
	for true {
		go kafkaConsumer.Poll(1, successEvent, errorEvent)
		select {
		case successResponse := <-successEvent:
			responseChan <- successResponse
		case errorResponse := <-errorEvent:
			errorChan <- errorResponse
		default:
			continue
		}
	}
}

// StartConsumer is the starting point to create a new consumer
func (c *Client) StartConsumer(newConsumer kafka.Consumer) bool {
	consumerConfig := newConsumer.GetConfig()
	errorChannel := make(chan error)
	triggerUnMarshalChan := make(chan wrapper.KafkaEventResponse)
	isEventProcessed := make(chan kafka.Consumer)
	jobs := make(chan kafka.Consumer, consumerConfig.WorkersCountToProcessData)
	subscriptionErr := c.KafkaConsumer.SubscribeTopics(consumerConfig.Topics, false)
	if subscriptionErr != nil {
		log.Fatalln("Unable to subscribe to Kafka Topics -> ", consumerConfig.Topics)
	}
	go listenAndTriggerUnmarshal(c.KafkaConsumer, triggerUnMarshalChan, errorChannel)
	go unmarshalAndNotifyJobsQueue(newConsumer, triggerUnMarshalChan, isEventProcessed, errorChannel)
	go pushEventToJobsQueue(jobs, isEventProcessed, errorChannel)
	for w := 1; w <= consumerConfig.WorkersCountToProcessData; w++ {
		go c.worker(jobs)
	}
	return true
}

/*
	worker will be for processing the the unmarshalled data. Each type of job have different processing
*/
func (c *Client) worker(jobs <-chan kafka.Consumer) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic: Recovered in defer -> ", r)
		}
	}()
	for j := range jobs {
		j.ProcessData()
		if !j.GetConfig().AutoCommit {
			// If auto-commit is set to false, it mean the offset had to be committed after message is read.
			topicPartitions, err := c.KafkaConsumer.Commit()
			if err != nil {
				for _, topicPartition := range topicPartitions {
					log.Println(fmt.Sprintf("Error: Error committing offset for topic %s with offset %v in parition %v",
						topicPartition.Topic, topicPartition.Offset, topicPartition.Partition))
				}
			}
		}
	}
}
