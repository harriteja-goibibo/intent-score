package grpc

import (
	"context"

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
}

// NewServer creates new grpc server. The context is used to close all
// all open resources created by it & its children.
func NewServer(ctx context.Context, cfg *ServerConfig) *Server {
	s := Server{
		config: cfg,
	}
	return &s
}

// PingPong returns "pong" response for the request.
func (s *Server) PingPong(ctx context.Context, request *health.HealthCheck) (response *health.HealthCheck, err error) {
	s.config.Logger.Info(ctx, "SERVER", "PingPong", "Obtained request", "")

	response = new(health.HealthCheck)

	response.Value = "pong"

	return response, nil
}
