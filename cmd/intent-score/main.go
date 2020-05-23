package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net"
	"runtime"
	"runtime/debug"

	"github.com/goibibo/intent-score/internal/config"
	serverImpl "github.com/goibibo/intent-score/internal/server/grpc"
	grpcAPI "github.com/goibibo/intent-score/pkg/api/grpc"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/goibibo/intent-score/internal/logger"
)

// initialization sets Golang configuration like GCFrequency & MaxProcs
func initialization() {
	// Garbage collector starts running when 70% of the memory is consumed.
	debug.SetGCPercent(70)

	// Number of cores which Go will use. Here if it's a 8 core machine.8 core is used.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func getConfig(path string) *viper.Viper {

	// Initialize the conf file with conf path.
	conf, err := config.New(path)

	if err != nil {
		panic(errors.Wrap(err, "Error in Config Initialisation"))
	}

	return conf
}

func createServerConfig(cfgPath string) (*serverImpl.ServerConfig, string) {

	conf := getConfig(cfgPath)

	// Initialise Loggers.
	serverLogger, err := logger.New("[gRPC-Server]", conf.GetString("settings.logger.info"))
	if err != nil {
		panic(errors.Wrap(err, "createServer error in creating loggers"))
	}

	// Attaching all resources to the server
	cfg := serverImpl.ServerConfig{
		Config: conf,
		Logger: serverLogger,
	}

	return &cfg, conf.GetString("settings.server.HOST")
}

func main() {

	initialization()

	cfgPath := flag.String("conf-path", "config/dev", "relative path to the conf directory")
	flag.Parse()

	srvCfg, port := createServerConfig(*cfgPath)

	// Initializing server with the main context.
	ctx, cancel := context.WithCancel(context.Background())
	server := serverImpl.NewServer(ctx, srvCfg)

	listen, err := net.Listen("tcp", port)

	if err != nil {
		panic(errors.Wrap(err, "gRPC connection is not started on the address "+port))
	}

	s := grpc.NewServer()
	grpcAPI.RegisterIntentScoreServer(s, server)

	// Channel returns if the application is successfully closed.
	closed := make(chan bool, 1)

	{
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

		go func() {
			// Go routine to read any stop signals.
			// Maintain the order of closer of the resouces.
			/*
				1) The tcp listener to the grpc server is stopped. This
				   makes sure that no more traffic is relayed to the server.
				2) The resources & routines created by the server i.e. the
				   grpc is closed.
				3) The created resources that were attached to the server should be closed.
			*/
			<-c
			signal.Stop(c)
			s.Stop()
			cancel()
			srvCfg.Logger.Close()
			closed <- true
		}()
	}

	fmt.Println("Server started at:", port, "!!!")
	if err := s.Serve(listen); err != nil {
		panic(errors.Wrap(err, "Error at grpc Serve"))
	}

	select {
	case <-closed:
		fmt.Println("Resources Successfully Closed :D")
	case <-time.After(5 * time.Second):
		fmt.Println("Closed by timeout :(")
	}
}
