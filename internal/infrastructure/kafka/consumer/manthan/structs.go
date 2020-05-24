package manthan

import (
	"github.com/goibibo/intent-score/internal/core"
	kafkaWrapper "github.com/goibibo/intent-score/internal/infrastructure/kafka"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"
)

type ManthanClient struct {
	Config kafkaWrapper.Config
	/*
		ProcessedEventData is the type of struct to which the raw bytes data from kafka will be un-marshalled to.
	*/
	ProcessedEventData *manthan.ManthanRealTimeData
	/*
		PromoUsage is the actual implementation of updating the usage of a promo code.
	*/
	ScoreSetter core.ScoreRepository
}
