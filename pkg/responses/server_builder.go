package responses

import "log"

type ServerResponseInterface interface {
	CreateServerResponse()
	WithInternalCode(code int)
	WithMessage(message interface{})
	BuildResponse() *ServerResponse
}

type ServerResponseBuilder struct {
	serverResponse *ServerResponse
}

func (srb *ServerResponseBuilder) CreateServerResponse() {
	nr := new(ServerResponse)
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

func (srb *ServerResponseBuilder) BuildResponse() *ServerResponse {
	if srb.serverResponse == nil {
		log.Fatal("ServerResponseBuilder should be initialized first before assign values.")
	}

	return srb.serverResponse
}
