package utils

import (
	"github.com/google/uuid"
	"github.com/space-devops/mountebank-sidecar/pkg/responses"
	"time"
)

func BuildApiResponse(internalCode int, message interface{}, correlationId string) *responses.Wrapper {
	sr := buildServerResponse(internalCode, message)
	return buildWrapperResponse(correlationId, sr)
}

func buildServerResponse(internalCode int, message interface{}) *responses.ServerResponse {
	srb := new(responses.ServerResponseBuilder)
	srb.CreateServerResponse()
	srb.WithInternalCode(internalCode)
	srb.WithMessage(message)
	return srb.BuildResponse()
}

func buildWrapperResponse(correlationId string, payload interface{}) *responses.Wrapper {
	wrb := new(responses.WrapperResponseBuilder)
	wrb.CreateWrapperResponse()
	wrb.WithCorrelationId(correlationId)
	wrb.WithTimestamp(time.Now().Format(time.RFC1123Z))
	wrb.WithPayload(payload)

	return wrb.BuildResponse()
}

func GenerateCorrelationId() string {
	return uuid.NewString()
}
