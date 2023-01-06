package config

type ServerConfig struct {
	Http HttpConfig `yaml:"http"`
	Grpc GrpcConfig `yaml:"grpc"`
}

type HttpConfig struct {
	Port                int `yaml:"port"`
	ReadTimeoutSeconds  int `yaml:"readTimeoutSeconds"`
	WriteTimeoutSeconds int `yaml:"writeTimeoutSeconds"`
}

type GrpcConfig struct {
	Port int `yaml:"port"`
}

type LoggerConfig struct {
	File  string `yaml:"file"`
	Level string `yaml:"level"`
}

type GlobalConfig struct {
	CorrelationIdHeader string `yaml:"correlationIdHeader"`
}

type MountebankConfig struct {
	Host      string            `yaml:"host"`
	Health    HealthConfig      `yaml:"health"`
	Imposters []ImpostersConfig `yaml:"imposters"`
}

type HealthConfig struct {
	Port int    `yaml:"port"`
	Path string `yaml:"path"`
}

type ImpostersConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

type Configurations struct {
	Global     GlobalConfig     `yaml:"global"`
	Server     ServerConfig     `yaml:"server"`
	Logger     LoggerConfig     `yaml:"logger"`
	Mountebank MountebankConfig `yaml:"mountebank"`
}
