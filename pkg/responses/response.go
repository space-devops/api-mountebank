package responses

type Wrapper struct {
	CorrelationId string      `json:"correlation_id"`
	Timestamp     string      `json:"timestamp"`
	Payload       interface{} `json:"payload"`
}

type ServerResponse struct {
	InternalCode int         `json:"internal_code"`
	Message      interface{} `json:"message"`
}
