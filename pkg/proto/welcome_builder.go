package space

import (
	"log"
)

// #################################################################################

type GrpcWelcomePayloadInterface interface {
	CreateWelcomePayload()
	WithInternalCode(code int)
	WithMessage(message string)
	BuildResponse() *WelcomePayload
}

type GrpcWelcomePayload struct {
	payload *WelcomePayload
}

func (wp *GrpcWelcomePayload) CreateWelcomePayload() {
	wp.payload = new(WelcomePayload)
}

func (wp *GrpcWelcomePayload) WithInternalCode(code int) {
	if wp.payload == nil {
		log.Fatal("GrpWelcomePayload should be initialized first before assign values.")
	}

	wp.payload.InternalCode = int32(code)
}

func (wp *GrpcWelcomePayload) WithMessage(message string) {
	wp.payload.Message = message
}

func (wp *GrpcWelcomePayload) BuildResponse() *WelcomePayload {
	return wp.payload
}

// #################################################################################

type GrpcWelcomeMessageInterface interface {
	CrateWelcomeMessage()
	WithCorrelationId(correlationId string)
	WithTimestamp(timestamp string)
	WithPayload(*WelcomePayload)
	BuildResponse() *WelcomeMessage
}

type GrpcWelcomeMessage struct {
	grpcWelcomeMessage *WelcomeMessage
}

func (wm *GrpcWelcomeMessage) CreateWelcomeMessage() {
	wm.grpcWelcomeMessage = new(WelcomeMessage)
}

func (wm *GrpcWelcomeMessage) WithCorrelationId(correlationId string) {
	if wm.grpcWelcomeMessage == nil {
		log.Fatal("GrpWelcomeMessage should be initialized first before assign values.")
	}

	wm.grpcWelcomeMessage.CorrelationId = correlationId
}

func (wm *GrpcWelcomeMessage) WithTimestamp(timestamp string) {
	if wm.grpcWelcomeMessage == nil {
		log.Fatal("GrpWelcomeMessage should be initialized first before assign values.")
	}

	wm.grpcWelcomeMessage.Timestamp = timestamp
}

func (wm *GrpcWelcomeMessage) WithPayload(payload *WelcomePayload) {
	if wm.grpcWelcomeMessage == nil {
		log.Fatal("GrpWelcomeMessage should be initialized first before assign values.")
	}

	wm.grpcWelcomeMessage.Payload = payload
}

func (wm *GrpcWelcomeMessage) BuildResponse() *WelcomeMessage {
	return wm.grpcWelcomeMessage
}
