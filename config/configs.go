package config

import (
	"fmt"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v3"
	"os"
	"sync"
)

const (
	YAML_PATH        = "infra/qbit-be.%s"
	ENV_MODE         = "ENV_MODE"
	DEFAULT_ENV_MODE = "development"
)

var (
	validEnvMode = map[string]struct{}{
		"local":       {},
		"development": {},
		"production":  {},
	}
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"postgresql"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Payment  Payment  `mapsstructure:"payment"`
}

var (
	config     *Config
	configOnce sync.Once
)

func LoadConfig() *Config {
	envMode := os.Getenv(ENV_MODE)
	if _, ok := validEnvMode[envMode]; !ok {
		envMode = DEFAULT_ENV_MODE // default env mode
	}

	cfgFilePath := fmt.Sprintf(YAML_PATH, envMode)

	configOnce.Do(func() {
		v := viper.New()
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.SetConfigName(cfgFilePath)
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %s", err))
		}

		config = &Config{}
		if err := v.Unmarshal(config); err != nil {
			panic(fmt.Errorf("failed to unmarshal config: %s", err))
		}
	})

	return config
}

func (c *Config) Auths() *AuthConfig {
	return &AuthConfig{
		jwtTokenSecret:     c.Jwt.Token.Secret,
		jwtTokenExpiresTTL: c.Jwt.Token.ExpiresTTL,
	}
}

func (c *Config) AuthPayment() *PaymentConfig {
	return &PaymentConfig{
		ApiKey:       c.Payment.ApiKey,
		MerchantCode: c.Payment.MerchantCode,
	}
}
