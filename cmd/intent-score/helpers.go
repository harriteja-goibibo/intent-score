package main

import (
	"github.com/goibibo/intent-score/internal"
	"github.com/goibibo/intent-score/internal/core/repository"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka/consumer"
	"log"
	"strconv"
	"strings"
	"time"

	infra "github.com/goibibo/intent-score/internal/infrastructure"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka/wrapper"

	"github.com/spf13/viper"

	"github.com/aerospike/aerospike-client-go"
	infraAero "github.com/goibibo/intent-score/internal/infrastructure/aerospike"
	"github.com/goibibo/intent-score/internal/infrastructure/kafka/consumer/manthan"
)

func newConsumerClient(kafkaConfig *viper.Viper, consumerGroupID string, autoCommit bool) *consumer.Client {
	consumerConfig := wrapper.ConsumerConfig{
		HeartBeatInterval:       kafkaConfig.GetInt("heart-beat-interval"),
		SessionTimeout:          kafkaConfig.GetInt("session-timeout"),
		MaxRetries:              kafkaConfig.GetInt("max-retries"),
		IntervalBetweenRetries:  kafkaConfig.GetInt("interval-between-retries"),
		MetadataRefreshInterval: kafkaConfig.GetInt("metadata-refresh-interval"),
		BootstrapServers:        strings.Join(kafkaConfig.GetStringSlice("bootstrap-servers"), ","),
		PartitionStrategy:       kafkaConfig.GetString("partition-strategy"),
		ConsumerGroupID:         consumerGroupID,
		AutoCommit:              autoCommit,
	}

	confluentKafkaClient, _ := wrapper.KafkaConsumer(consumerConfig)
	/*
		A wrapper client on top of actual kafka functionality.
	*/
	kafkaConsumerWrapper, _ := wrapper.NewConsumerClient(confluentKafkaClient)
	/*
		Our custom consumer where all application code resides.
		this deals with the wrapper we have created in previous step.
	*/
	kafkaClient := consumer.New(kafkaConsumerWrapper)
	return kafkaClient
}

// newAerospikeClient creates a new Aerospike Client.
func newAerospikeClient(aerospikeConfig *viper.Viper) infra.Aerospike {
	hosts := make([]*aerospike.Host, 0)
	for _, hostConfig := range aerospikeConfig.GetStringSlice("hosts") {
		hostSplits := strings.SplitN(hostConfig, ":", 2)
		hostName := hostSplits[0]
		port, _ := strconv.Atoi(hostSplits[1])
		hosts = append(hosts, &aerospike.Host{Name: hostName, Port: port})
	}

	aerospikeClientPolicy := aerospike.NewClientPolicy()
	aerospikeClientPolicy.ConnectionQueueSize = aerospikeConfig.GetInt("connection-queue-size")
	aerospikeClientPolicy.LimitConnectionsToQueueSize = aerospikeConfig.GetBool("limit-connections-queue-size")
	aerospikeClientPolicy.Timeout = 1000 * time.Millisecond

	if len(hosts) == 0 {
		log.Fatalf("No Aerospike Host Found Config: %+v", aerospikeConfig.AllSettings())
	}

	client, err := aerospike.NewClientWithPolicyAndHost(aerospikeClientPolicy, hosts...)
	if err != nil {
		log.Fatalf("NewClientWithPolicyAndHost: Error in Connection Error: %s", err.Error())
	}

	aeroClient, err := infraAero.NewAeroClient(client)
	if err != nil {
		log.Fatalf("NewAeroClient: Error in creating Infra's Aerospike Error: %s", err.Error())
	}

	return aeroClient
}

func consumeManthanRealTimeData(realTimeData *viper.Viper, aeroClient infra.Aerospike, logger internal.Logger) {

	if !realTimeData.GetBool("enable") {
		log.Println("Manthan RealTimeData Consumer has been Disabled !!!")
		return
	}

	log.Println("Manthan RealTimeData Consumer has started :)")

	consumerGroupID := realTimeData.GetString("consumer-group-id")
	autoCommit := realTimeData.GetBool("auto-commit")

	realTimeDataClient := newConsumerClient(realTimeData, consumerGroupID, autoCommit)
	scoreRepo, _ := repository.NewScoreRepo(aeroClient)
	realTimeDataConsumer := manthan.NewManthanDataClient(kafka.Config{
		Topics:                    realTimeData.GetStringSlice("topics"),
		WorkersCountToProcessData: realTimeData.GetInt("worker-count"),
		LogBucket:                 realTimeData.GetString("log-bucket"),
		AutoCommit:                autoCommit,
	}, scoreRepo)

	realTimeDataClient.StartConsumer(realTimeDataConsumer)
}

func startKafkaConsumers(config *viper.Viper, aeroClient infra.Aerospike, logger internal.Logger) {

	if !config.GetBool("enable") {
		log.Println("Kafka Consumer has been Disabled !!")
		return
	}

	log.Println("Starting Kafka Consumer")
	{
		if config.InConfig("manthan.real_time_data") {
			consumeManthanRealTimeData(config.Sub("manthan.real_time_data"), aeroClient, logger)
		}
	}
}
