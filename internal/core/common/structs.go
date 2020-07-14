package common

import (
	"github.com/goibibo/intent-score/pkg/api/common"
)

const (
	TRAVEL_DATE_FACTOR = 10
)

type RealTimeData struct {
	Org         common.Org
	Vertical    common.Vertical
	PageHitType common.PageHitType
	RequestDate int
	TravelDate  int
	UserId      string
	EntityId    string
	RoomNights  int
	UserHistory
}

type UserHistory struct {
	// Advanced Purchase Min --> Diff of bookingData and TravelData
	// Advanced Purchase Max --> Diff of bookingData and TravelData
	// Indicated the review page, payment page bounce rate
	// UpdatedTs When the transaction is updated
	ApMin       int
	ApMax       int
	Sensitivity int
	UpdatedTs   string
}

type EntityData struct {
	Score      int
	TodayScore int
	UpdatedTs  string
}

type RealTimeDataResponse struct {
	RequestDate int
	EntityId    string
	TravelDate  int
	RoomNights  int
	PageHits    []int
}

type UserData struct {
	RealTimeData []RealTimeDataResponse
	EntityData   EntityData
}
