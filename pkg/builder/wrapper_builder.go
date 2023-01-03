package builder

import (
	"github.com/space-devops/api-mountebank/pkg/responses"
	"log"
)

type WrapperResponseBuilderInterface interface {
	CreateWrapperResponse()
	WithCorrelationId(correlationId string)
	WithTimestamp(timestamp string)
	WithPayload(payload interface{})
	BuildResponse() *responses.Wrapper
}

type WrapperResponseBuilder struct {
	wrapper *responses.Wrapper
}

func (wrb *WrapperResponseBuilder) CreateWrapperResponse() {
	wr := new(responses.Wrapper)
	wrb.wrapper = wr
}

func (wrb *WrapperResponseBuilder) WithCorrelationId(correlationId string) {
	if wrb.wrapper == nil {
		log.Fatal("WrapperResponseBuilder should be initialized first before assign values.")
	}
	wrb.wrapper.CorrelationId = correlationId
}

func (wrb *WrapperResponseBuilder) WithTimestamp(timestamp string) {
	if wrb.wrapper == nil {
		log.Fatal("WrapperResponseBuilder should be initialized first before assign values.")
	}
	wrb.wrapper.Timestamp = timestamp
}

func (wrb *WrapperResponseBuilder) WithPayload(payload interface{}) {
	if wrb.wrapper == nil {
		log.Fatal("WrapperResponseBuilder should be initialized first before assign values.")
	}
	wrb.wrapper.Payload = payload
}

func (wrb *WrapperResponseBuilder) BuildResponse() *responses.Wrapper {
	if wrb.wrapper == nil {
		log.Fatal("WrapperResponseBuilder should be initialized first before assign values.")
	}
	return wrb.wrapper
}
