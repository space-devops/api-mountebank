package config

type ServerConfig struct {
	Port                int `yaml:"port"`
	ReadTimeoutSeconds  int `yaml:"readTimeoutSeconds"`
	WriteTimeoutSeconds int `yaml:"writeTimeoutSeconds"`
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
	Imposters []ImpostersConfig `yaml:"imposters"`
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
