package config

import (
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"github.com/spf13/viper"
)

var cfg *Configurations

func GetConfig() *Configurations {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {
	cfg = readConfigFile()
}

func readConfigFile() *Configurations {
	viper.SetConfigName(utils.ConfigFileName)
	viper.SetConfigType(utils.ConfigFileExtension)

	for _, path := range utils.ConfigFilePaths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setConfigDefaults()
		} else {
			logger.LogPanic("Error while reading configuration from file", utils.NoCorrelationId)
		}
	}

	var globalConfig Configurations

	if err := viper.Unmarshal(&globalConfig); err != nil {
		logger.LogPanic("Error while parsing configuration from file", utils.NoCorrelationId)
	}

	return &globalConfig
}

func setConfigDefaults() {
	viper.SetDefault("server.port", utils.ServerPort)
	viper.SetDefault("server.readTimeoutSeconds", utils.ServerReadTimeout)
	viper.SetDefault("server.writeTimeoutSeconds", utils.ServerWriteTimeout)

	viper.SetDefault("logger.file", utils.LoggerFileName)
	viper.SetDefault("logger.level", utils.LoggerLevel)

	viper.SetDefault("global.correlationIdHeader", utils.CorrelationIdHeaderName)
}
