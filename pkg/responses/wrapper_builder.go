package responses

import "log"

type WrapperResponseBuilderInterface interface {
	CreateWrapperResponse()
	WithCorrelationId(correlationId int)
	WithTimestamp(timestamp string)
	WithPayload(payload interface{})
	BuildResponse() *Wrapper
}

type WrapperResponseBuilder struct {
	wrapper *Wrapper
}

func (wrb *WrapperResponseBuilder) CreateWrapperResponse() {
	wr := new(Wrapper)
	wrb.wrapper = wr
}

func (wrb *WrapperResponseBuilder) WithCorrelationId(correlationId int) {
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

func (wrb *WrapperResponseBuilder) BuildResponse() *Wrapper {
	if wrb.wrapper == nil {
		log.Fatal("WrapperResponseBuilder should be initialized first before assign values.")
	}
	return wrb.wrapper
}
