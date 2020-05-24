package grpc

import (
	"context"
	"github.com/goibibo/intent-score/internal/core"
	"github.com/goibibo/intent-score/internal/core/repository"
	"github.com/goibibo/intent-score/internal/impls/real_time_data"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"

	"github.com/goibibo/intent-score/internal"
	infra "github.com/goibibo/intent-score/internal/infrastructure"
	"github.com/goibibo/intent-score/pkg/api/grpc/health"

	"github.com/spf13/viper"
)

// ServerConfig has the configurations required
// for the functioning of the server.
type ServerConfig struct {
	Config    *viper.Viper
	Logger    internal.Logger
	Aerospike infra.Aerospike
}

// Server is the main struct used for gRPC server.
type Server struct {
	config *ServerConfig

	repo struct {
		ScoreRepo core.ScoreRepository
	}
}

// NewServer creates new grpc server. The context is used to close all
// all open resources created by it & its children.
func NewServer(ctx context.Context, cfg *ServerConfig) *Server {
	s := Server{
		config: cfg,
	}
	s.repo.ScoreRepo, _ = repository.NewScoreRepo(s.config.Aerospike)
	return &s
}

// PingPong returns "pong" response for the request.
func (s *Server) PingPong(ctx context.Context, request *health.HealthCheck) (response *health.HealthCheck, err error) {
	s.config.Logger.Info(ctx, "SERVER", "PingPong", "Obtained request", "")

	response = new(health.HealthCheck)

	response.Value = "pong"

	return response, nil
}

func (s *Server) ManthanRealTimeData(ctx context.Context, request *manthan.ManthanRealTimeData) (response *manthan.ManthanRealTimeData, err error) {
	err = real_time_data.SaveRealTimeData(ctx, request, s.repo.ScoreRepo)
	return new(manthan.ManthanRealTimeData), err
}
