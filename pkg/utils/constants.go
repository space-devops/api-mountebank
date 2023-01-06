package utils

const (
	ServerPort              = 2000
	ServerWriteTimeout      = 15
	ServerReadTimeout       = 15
	GrpcServerPort          = 6000
	CorrelationIdHeaderName = "X-Internal-Correlation-ID"
	NoCorrelationId         = "global"
	LoggerFileName          = "default.log"
	LoggerLevel             = "debug"
	DefaultContentType      = "application/json"

	ConfigFileName      = "config"
	SecretFileName      = "secrets"
	ConfigFileExtension = "yaml"
	SecretFileExtension = "toml"
)

var (
	ConfigFilePaths = [4]string{"./config", "$HOME/.imposters", ".", "$HOME/.creds"}
)
