package manthan

import (
	"context"
	"encoding/json"
	"github.com/goibibo/intent-score/internal/core"
	"github.com/goibibo/intent-score/internal/impls/real_time_data"
	kafkaWrapper "github.com/goibibo/intent-score/internal/infrastructure/kafka"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"
	"github.com/pkg/errors"
)

func NewManthanDataClient(config kafkaWrapper.Config, scoreRepo core.ScoreRepository) *ManthanClient {
	return &ManthanClient{
		ScoreSetter:        scoreRepo,
		ProcessedEventData: new(manthan.ManthanRealTimeData),
		Config:             config,
	}
}

func (c *ManthanClient) GetConfig() kafkaWrapper.Config {
	return c.Config
}
func (c *ManthanClient) Unmarshal(data map[string][]byte) (kafkaWrapper.Consumer, error) {
	processedData := new(manthan.ManthanRealTimeData)

	err := json.Unmarshal(data["data"], processedData)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to unmarshal kafka event data with the specified error")
	}
	c.ProcessedEventData = processedData
	return c, nil
}

func (c *ManthanClient) ProcessData() {
	if c == nil {
		// TODO error handling
		return
	}
	processedData := c.ProcessedEventData

	err := real_time_data.SaveRealTimeData(context.Background(), processedData, c.ScoreSetter)
	if err != nil {
		// TODO : Handle error
	}

}
