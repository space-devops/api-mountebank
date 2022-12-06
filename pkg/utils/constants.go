package utils

import "time"

const (
	ServerPort              = 3000
	ServerWriteTimeout      = 15 * time.Second
	ServerReadTimeout       = 15 * time.Second
	CorrelationIdHeaderName = "X-Internal-Correlation-ID"
)
