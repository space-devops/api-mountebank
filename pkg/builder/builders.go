package builder

import (
	"github.com/space-devops/api-mountebank/pkg/responses"
	"time"
)

func BuildApiResponse(internalCode int, message interface{}, correlationId string) *responses.Wrapper {
	sr := buildServerResponse(internalCode, message)
	return buildWrapperResponse(correlationId, sr)
}

func buildServerResponse(internalCode int, message interface{}) *responses.ServerResponse {
	srb := new(ServerResponseBuilder)
	srb.CreateServerResponse()
	srb.WithInternalCode(internalCode)
	srb.WithMessage(message)
	return srb.BuildResponse()
}

func buildWrapperResponse(correlationId string, payload interface{}) *responses.Wrapper {
	wrb := new(WrapperResponseBuilder)
	wrb.CreateWrapperResponse()
	wrb.WithCorrelationId(correlationId)
	wrb.WithTimestamp(time.Now().Format(time.RFC1123Z))
	wrb.WithPayload(payload)

	return wrb.BuildResponse()
}
