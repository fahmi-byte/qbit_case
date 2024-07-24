package config

type Jwt struct {
	Token Token `mapstructure:"token"`
}

type Token struct {
	ExpiresTTL int    `mapstructure:"expires-ttl"`
	Secret     string `mapstructure:"secret"`
}
