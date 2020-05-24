package common

import (
	"github.com/goibibo/intent-score/pkg/api/common"
)

type RealTimeData struct {
	Org         common.Org
	Vertical    common.Vertical
	PageHitType common.PageHitType
	RequestDate int64
	UserId      string
	EntityId    string
	TravelDate  string
	RoomNights  int
}
