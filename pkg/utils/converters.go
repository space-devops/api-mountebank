package utils

import (
	"encoding/json"
	"github.com/space-devops/api-mountebank/pkg/logger"

	"time"
)

func IntToSeconds(number int) time.Duration {
	return time.Duration(number) * time.Second
}

func ObjectToJsonObject(raw interface{}, cid string) ([]byte, error) {
	jr, je := json.Marshal(raw)
	if je != nil {
		logger.LogError("Error marshalling responses", cid, logger.LogExtraInfo{
			Key:   "MarshallerError",
			Value: je.Error(),
		})
		return nil, je
	}

	return jr, nil
}

func JsonObjectToObject(raw []byte, dest interface{}, cid string) error {
	je := json.Unmarshal(raw, dest)
	if je != nil {
		logger.LogError("Error marshalling responses", cid, logger.LogExtraInfo{
			Key:   "MarshallerError",
			Value: je.Error(),
		})
		return je
	}

	return nil
}
