package repository

import (
	"context"
	"fmt"

	"github.com/goibibo/intent-score/internal/core/common"
	infra "github.com/goibibo/intent-score/internal/infrastructure"
)

type ScoreRepo struct {
	client infra.Aerospike
}

// NewPromoRepo creates the repository of fetcher module on top of the provided aerospike client.
func NewPromoRepo(client infra.Aerospike) (*ScoreRepo, error) {
	f := ScoreRepo{
		client: client,
	}
	return &f, nil
}

func (s *ScoreRepo) SaveManthanRealTimeData(ctx context.Context, data common.RealTimeData) error {
	fmt.Println("data here ----1")
	fmt.Println("data here ------2")
	fmt.Println("data here -------3")
	return nil
}
