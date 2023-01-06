package config

import (
	"fmt"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"github.com/spf13/viper"
)

var cfg *Configurations
var scr *SecretConfig

func GetConfig() *Configurations {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func GetSecrets() *SecretConfig {
	if scr == nil {
		loadSecrets()
	}

	return scr
}

func loadConfig() {
	vcfg := viper.New()
	cfg = readConfigFile(utils.ConfigFileName, utils.ConfigFileExtension, vcfg)
}

func loadSecrets() {
	vscr := viper.New()
	scr = readSecretFile(utils.SecretFileName, utils.SecretFileExtension, vscr)
}

func readConfigFile(fileName string, fileExtension string, vp *viper.Viper) *Configurations {
	if err := readFile(fileName, fileExtension, vp); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setConfigDefaults(vp)
		} else {
			logger.LogPanic("Error while reading configuration from file", utils.NoCorrelationId)
		}
	}

	var globalConfig Configurations
	unmarshalConfig(&globalConfig, vp)
	fmt.Printf("Config file: %v\n", globalConfig)

	return &globalConfig
}

func readSecretFile(fileName string, fileExtension string, vp *viper.Viper) *SecretConfig {
	if err := readFile(fileName, fileExtension, vp); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			setSecretDefaults(vp)
		} else {
			logger.LogPanic("Error while reading secrets from file", utils.NoCorrelationId)
		}
	}

	var globalSecrets SecretConfig
	unmarshalConfig(&globalSecrets, vp)
	fmt.Printf("Config file: %v\n", globalSecrets)

	return &globalSecrets
}

func readFile(fileName string, fileExtension string, vp *viper.Viper) error {
	vp.SetConfigName(fileName)
	vp.SetConfigType(fileExtension)

	for _, path := range utils.ConfigFilePaths {
		vp.AddConfigPath(path)
	}

	return vp.ReadInConfig()
}

func unmarshalConfig(object interface{}, vp *viper.Viper) {
	if err := vp.Unmarshal(object); err != nil {
		logger.LogPanic("Error while parsing object from file", utils.NoCorrelationId)
	}
}

func setConfigDefaults(vp *viper.Viper) {
	vp.SetDefault("server.http.port", utils.ServerPort)
	vp.SetDefault("server.http.readTimeoutSeconds", utils.ServerReadTimeout)
	vp.SetDefault("server.http.writeTimeoutSeconds", utils.ServerWriteTimeout)

	vp.SetDefault("server.grpc.port", utils.GrpcServerPort)

	vp.SetDefault("logger.file", utils.LoggerFileName)
	vp.SetDefault("logger.level", utils.LoggerLevel)

	vp.SetDefault("global.correlationIdHeader", utils.CorrelationIdHeaderName)
}

func setSecretDefaults(vp *viper.Viper) {
	vp.SetDefault("secrets.enable", false)
	vp.SetDefault("secrets.db.username", "defaultDbName")
	vp.SetDefault("secrets.db.password", "defaultDbPassword")
}
