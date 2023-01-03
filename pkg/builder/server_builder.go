package builder

import (
	"github.com/space-devops/api-mountebank/pkg/responses"
	"log"
)

type ServerResponseInterface interface {
	CreateServerResponse()
	WithInternalCode(code int)
	WithMessage(message interface{})
	BuildResponse() *responses.ServerResponse
}

type ServerResponseBuilder struct {
	serverResponse *responses.ServerResponse
}

func (srb *ServerResponseBuilder) CreateServerResponse() {
	nr := new(responses.ServerResponse)
	srb.serverResponse = nr
}

func (srb *ServerResponseBuilder) WithInternalCode(code int) {
	if srb.serverResponse == nil {
		log.Fatal("ServerResponseBuilder should be initialized first before assign values.")
	}

	srb.serverResponse.InternalCode = code
}

func (srb *ServerResponseBuilder) WithMessage(message interface{}) {
	if srb.serverResponse == nil {
		log.Fatal("ServerResponseBuilder should be initialized first before assign values.")
	}

	srb.serverResponse.Message = message
}

func (srb *ServerResponseBuilder) BuildResponse() *responses.ServerResponse {
	if srb.serverResponse == nil {
		log.Fatal("ServerResponseBuilder should be initialized first before assign values.")
	}

	return srb.serverResponse
}
