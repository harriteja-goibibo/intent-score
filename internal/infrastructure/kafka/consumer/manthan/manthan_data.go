package manthan

import (
	"context"
	"encoding/json"
	"github.com/goibibo/intent-score/internal/core"
	"github.com/goibibo/intent-score/internal/core/common"
	"github.com/goibibo/intent-score/internal/helpers/time"
	kafkaWrapper "github.com/goibibo/intent-score/internal/infrastructure/kafka"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"
	"strconv"

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

	realTimeData := common.RealTimeData{
		Org:         processedData.GetOrg(),
		Vertical:    processedData.GetVertical(),
		PageHitType: processedData.GetPageType(),
		UserId:      processedData.GetUserId(),
		EntityId:    processedData.GetEntityId(),
		TravelDate:  time.GetRequiredTimeFormat(processedData.GetTravelStartDate(), time.DateInYYYYMMDDFormat, time.DateInYYYYMMDDFormat),
	}

	requestDate, err := strconv.Atoi(time.GetRequiredTimeFormat(processedData.GetRequestDate(), time.DateInYYYYMMDDHHMMSSFormat, time.DateInYYYYMMDDFormat))
	if err != nil {
		realTimeData.RequestDate = int64(requestDate)
	}
	// TODO error handling
	travelStartDate, _ := strconv.Atoi(time.GetRequiredTimeFormat(processedData.GetTravelStartDate(), time.DateInYYYYMMDDHHMMSSFormat, time.DateInYYYYMMDDFormat))
	travelEndDate, _ := strconv.Atoi(time.GetRequiredTimeFormat(processedData.GetTravelEndDate(), time.DateInYYYYMMDDHHMMSSFormat, time.DateInYYYYMMDDFormat))
	realTimeData.RoomNights = travelStartDate - travelEndDate
	err = c.ScoreSetter.SaveManthanRealTimeData(context.Background(), realTimeData)
	if err != nil {
		// TODO : Handle error
	}

}
