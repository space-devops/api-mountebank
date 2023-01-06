package config

// TODO: We marked each config as json only to show each config as json in the api response

type SecretConfig struct {
	Secrets SecretConfigDetails `json:"secrets"`
}

type SecretConfigDetails struct {
	Enable bool            `json:"enable"`
	Db     DatabaseSecrets `json:"db"`
	Apis   []ApiSecrets    `json:"apis"`
}

type DatabaseSecrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ApiSecrets struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}
