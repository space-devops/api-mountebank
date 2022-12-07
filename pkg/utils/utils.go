package utils

import (
	"github.com/google/uuid"
)

func GenerateCorrelationId() string {
	return uuid.NewString()
}
