package core

import (
	"context"
	"github.com/goibibo/intent-score/internal/core/common"
)

type ScoreRepository interface {
	SaveManthanRealTimeData(ctx context.Context, data common.RealTimeData) error
}
