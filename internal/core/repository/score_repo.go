package repository

import (
	"context"
	"github.com/goibibo/intent-score/internal/core/common"
	"github.com/goibibo/intent-score/internal/helpers/time"
	infra "github.com/goibibo/intent-score/internal/infrastructure"
	commonPb "github.com/goibibo/intent-score/pkg/api/common"
	"strings"
)

type ScoreRepo struct {
	client infra.Aerospike
}

// NewScoreRepo creates the repository of fetcher module on top of the provided aerospike client.
func NewScoreRepo(client infra.Aerospike) (*ScoreRepo, error) {
	f := ScoreRepo{
		client: client,
	}
	return &f, nil
}

func getEntityDataBin(vertical string) string {
	return getVerticalDataBin(vertical) + aeroEntityDataBin
}

func getVerticalDataBin(vertical string) string {
	return strings.ToLower(vertical)
}

func (s *ScoreRepo) GetUserData(ctx context.Context, userId string, vertical string) (userHistory common.UserHistory, err error) {
	userHistory = common.UserHistory{}
	aeroKey := infra.Key{
		Namespace: aeroUserNamespace,
		SetName:   aeroRealTimeDataSetName,
		UserKey:   userId,
	}
	aeroData, _ := s.client.MapGetByKey(ctx, aeroKey, aeroUserDataBin, strings.ToLower(vertical))
	if aeroData != nil {
		data := aeroData.([]interface{})
		userHistory.ApMin = data[0].(int)
		userHistory.ApMax = data[1].(int)
		// TODO: Calculate the sensitivity factor from realTimeData
		userHistory.Sensitivity = data[2].(int)
		userHistory.UpdatedTs = data[3].(string)
	}
	return
}

func (s *ScoreRepo) SetUserData(ctx context.Context, userId string, vertical string, userHistory common.UserHistory) (err error) {
	aeroKey := infra.Key{
		Namespace: aeroUserNamespace,
		SetName:   aeroRealTimeDataSetName,
		UserKey:   userId,
	}
	err = s.client.MapPutByKey(ctx, aeroKey, aeroUserDataBin, strings.ToLower(vertical), []interface{}{userHistory.ApMin, userHistory.ApMax, userHistory.Sensitivity, userHistory.UpdatedTs})
	return
}

func (s *ScoreRepo) SaveRealTimeData(ctx context.Context, data common.RealTimeData) (userData common.UserData, err error) {
	userData = common.UserData{
		RealTimeData: make([]common.RealTimeDataResponse, 0),
		EntityData:   common.EntityData{},
	}
	aeroKey := infra.Key{
		Namespace: aeroUserNamespace,
		SetName:   aeroRealTimeDataSetName,
		UserKey:   data.UserId,
	}
	beginValue := []interface{}{time.SubstractDays(data.RequestDate, time.DateInYYYYMMDDFormat, data.UserHistory.ApMax+1), data.EntityId, time.SubstractDays(data.TravelDate, time.DateInYYYYMMDDFormat, common.TRAVEL_DATE_FACTOR), "*", "*"}
	endValue := []interface{}{time.AddDays(data.RequestDate, time.DateInYYYYMMDDFormat, 1), data.EntityId, time.AddDays(data.TravelDate, time.DateInYYYYMMDDFormat, common.TRAVEL_DATE_FACTOR), "*", "*"}
	aeroDataList, _ := s.client.OrderedListGetByValueRange(ctx, aeroKey, getVerticalDataBin(data.Vertical.String()), beginValue, endValue)
	var removeAeroData interface{}
	// SEARCH, DETAIL, REVIEW, PAYMENT, TRANSACTION Page Hits
	pageHitsData := []interface{}{0, 0, 0, 0, 0}
	pageHitsData[commonPb.PageHitType_value[data.PageHitType.String()]-1] = 1

	if aeroDataList != nil {
		for _, aeroDataInt := range aeroDataList.([]interface{}) {
			if aeroDataInt == nil {
				continue
			}
			aeroData := aeroDataInt.([]interface{})
			requestDate := aeroData[0].(int)
			travelDate := aeroData[2].(int)
			hitsData := aeroData[4].([]interface{})
			if requestDate == data.RequestDate && travelDate == data.TravelDate {
				removeAeroData = aeroDataInt
				_, err = s.client.OrderedListRemoveByValueList(ctx, aeroKey, getVerticalDataBin(data.Vertical.String()), []interface{}{removeAeroData})
				pageHitsData = hitsData
				pageHitsData[commonPb.PageHitType_value[data.PageHitType.String()]-1] = pageHitsData[commonPb.PageHitType_value[data.PageHitType.String()]-1].(int) + 1
			}
			userData.RealTimeData = append(userData.RealTimeData, common.RealTimeDataResponse{
				RequestDate: requestDate,
				EntityId:    data.EntityId,
				TravelDate:  travelDate,
				RoomNights:  aeroData[3].(int),
				PageHits:    []int{hitsData[0].(int), hitsData[1].(int), hitsData[2].(int), hitsData[3].(int), hitsData[4].(int)},
			})
		}
	}

	beginValue = []interface{}{[]interface{}{data.RequestDate, data.EntityId, data.TravelDate, data.RoomNights, pageHitsData}}
	err = s.client.OrderedListAppend(ctx, aeroKey, getVerticalDataBin(data.Vertical.String()), beginValue)

	// TODO: Handle Error
	aeroData, _ := s.client.MapGetByKey(ctx, aeroKey, getEntityDataBin(data.Vertical.String()), data.EntityId)
	if aeroData != nil {
		entityData := aeroData.([]interface{})
		userData.EntityData.Score = entityData[0].(int)
		userData.EntityData.TodayScore = entityData[1].(int)
		// TODO: time.Now() - updatedTs > 1, then club score and today's score
		userData.EntityData.UpdatedTs = entityData[3].(string)
	}

	return userData, err
}
