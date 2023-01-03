package space

import (
	"context"
	"github.com/space-devops/api-mountebank/pkg/client"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"time"
)

type PlanetServer struct{}

func (*PlanetServer) GetWelcome(context.Context, *emptypb.Empty) (*WelcomeMessage, error) {
	rsp := BuildGrpcWelcomeMessage(
		utils.GenerateCorrelationId(),
		time.Now().Format(time.RFC1123Z),
		http.StatusOK,
		"Hello world from GRPC Server",
	)

	return rsp, nil
}

func (*PlanetServer) GetPlanetList(context.Context, *emptypb.Empty) (*PlanetList, error) {
	cid := utils.GenerateCorrelationId()

	surl := client.GetServiceURL("list")
	raw, err := client.CallService(http.MethodGet, surl, cid)
	if err != nil {
		logger.LogPanic("unable to call backend service", cid)
	}

	return BuildPlanetList(raw, cid)
}

func (*PlanetServer) GetPlanetDetails(ctx context.Context, param *wrapperspb.StringValue) (*PlanetDetails, error) {
	cid := utils.GenerateCorrelationId()

	logger.LogInfo(param.Value, cid)
	surl := client.GetServiceURL(param.Value)
	raw, err := client.CallService(http.MethodGet, surl, cid)
	if err != nil {
		logger.LogPanic("unable to call backend service", cid)
	}

	return BuildPlanetDetails(raw, cid)
}

func (*PlanetServer) mustEmbedUnimplementedPlanetServiceServer() {
	logger.LogInfo("mustEmbedUnimplementedPlanetServiceServer", "GRPC Server")
}
