package space

import (
	"github.com/space-devops/api-mountebank/pkg/config"
	"time"
)

func BuildSecrets(secrets *config.SecretConfig, internalCode int, correlationId string) *SecretMessage {
	var gapis []*ApisSecrets

	for _, api := range secrets.Secrets.Apis {
		gapi := ApisSecrets{
			Host:     api.Host,
			Username: api.Username,
			Password: api.Password,
		}

		gapis = append(gapis, &gapi)
	}

	gdb := DatabaseSecrets{
		Username: secrets.Secrets.Db.Username,
		Password: secrets.Secrets.Db.Password,
	}

	sd := SecretDetails{
		Enable: secrets.Secrets.Enable,
		Db:     &gdb,
		Apis:   gapis,
	}

	s := Secrets{
		Secrets: &sd,
	}

	payload := SecretPayload{
		InternalCode: int32(internalCode),
		Message:      &s,
	}

	return &SecretMessage{
		CorrelationId: correlationId,
		Timestamp:     time.Now().Format(time.RFC1123Z),
		Payload:       &payload,
	}
}
