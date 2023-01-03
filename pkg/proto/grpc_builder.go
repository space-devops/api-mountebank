package space

func BuildGrpcWelcomeMessage(correlationId string, timestamp string, internalCode int, message string) *WelcomeMessage {
	payload := buildGrpcWelcomePayload(internalCode, message)

	gwm := new(GrpcWelcomeMessage)
	gwm.CreateWelcomeMessage()
	gwm.WithCorrelationId(correlationId)
	gwm.WithTimestamp(timestamp)
	gwm.WithPayload(payload)

	return gwm.BuildResponse()
}

func buildGrpcWelcomePayload(internalCode int, message string) *WelcomePayload {
	gwp := new(GrpcWelcomePayload)
	gwp.CreateWelcomePayload()
	gwp.WithInternalCode(internalCode)
	gwp.WithMessage(message)

	return gwp.BuildResponse()
}
