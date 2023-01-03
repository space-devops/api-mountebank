package grpc

import (
	"fmt"
	"github.com/space-devops/api-mountebank/pkg/config"
	"github.com/space-devops/api-mountebank/pkg/logger"
	space "github.com/space-devops/api-mountebank/pkg/proto"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"google.golang.org/grpc"
	"net"
)

func StartGRPC() {
	logger.LogInfo("GRPC Server starting...", utils.NoCorrelationId)

	// 1) Create listener
	listener, err := net.Listen("tcp", serverAddress())
	if err != nil {
		logger.LogPanic(err.Error(), utils.NoCorrelationId)
	}

	// 2) Create Server Instance
	server := grpc.NewServer()

	// 3) Register the server and the struct implementing the service interface
	space.RegisterPlanetServiceServer(server, new(space.PlanetServer))

	// 4) Register server reflection in order to access it from Evans CLI [OPTIONAL]

	// 5) Start server instance
	if err = server.Serve(listener); err != nil {
		logger.LogPanic(err.Error(), utils.NoCorrelationId)
	}
}

func serverAddress() string {
	grpcPort := config.GetConfig().Server.Grpc.Port
	return fmt.Sprintf("0.0.0.0:%d", grpcPort)
}
