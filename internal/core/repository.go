package core

import (
	"context"
	"github.com/goibibo/intent-score/internal/core/common"
)

type ScoreRepository interface {
	GetUserData(ctx context.Context, userId string, vertical string) (common.UserHistory, error)
	SetUserData(ctx context.Context, userId string, vertical string, userHistory common.UserHistory) (err error)
	SaveRealTimeData(ctx context.Context, data common.RealTimeData) (common.UserData, error)
}
