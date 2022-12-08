package utils

const (
	ServerPort              = 2000
	ServerWriteTimeout      = 15
	ServerReadTimeout       = 15
	CorrelationIdHeaderName = "X-Internal-Correlation-ID"
	NoCorrelationId         = "global"
	LoggerFileName          = "default.log"
	LoggerLevel             = "debug"

	ConfigFileName      = "config"
	ConfigFileExtension = "yaml"
)

var (
	ConfigFilePaths = [3]string{"./config", "$HOME/.imposters", "."}
)
