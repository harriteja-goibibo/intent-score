package kafka

// ConsumerConfig is the configuration of consumer that describes topic, groupID etc
type Config struct {
	// Topics that this consumer had to listen to.
	Topics []string
	// WorkersCountToProcessData is the count of number of workers which does the processing of kafka event in parallel
	WorkersCountToProcessData int
	// LogBucket defines in case if some info or error to be logged, this specifies which bucket the logs had to be logged to.
	LogBucket string
	// AutoCommit if set to true, as soon as event is received the offset will be committed. If not we commit manually.
	AutoCommit bool
}

type Consumer interface {
	// GetConfig Returns the individual Consumer Config that allows base to listen to a particular broker's topics.
	GetConfig() Config
	// Unmarshal transforms the data from bytes format to a format that the consumer's process function understands
	Unmarshal(map[string][]byte) (Consumer, error)
	// ProcessData is where all the application logic will be done.
	ProcessData()
}

type ListenToEvents interface {
	StartConsumer(newConsumer Consumer) bool
}
