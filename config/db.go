package config

import (
	"fmt"
	"time"
)

type Database struct {
	Host                  string `mapstructure:"host"`
	Port                  string `mapstructure:"port"`
	DB                    string `mapstructure:"db"`
	Driver                string `mapstructure:"driver"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	SslMode               string `mapstructure:"ssl-mode"`
	Debug                 bool   `mapstructure:"debug"`
	MaxIdleConnection     int    `mapstructure:"max-idle-connections"`
	MaxOpenConnections    int    `mapstructure:"max-open-connections"`
	ConnectionMaxLifetime int    `mapstructure:"connection-max-life-time"`
	ConnectionMaxIdleTime int    `mapstructure:"connection-max-idle-time"`
}

func (c Database) GetDataSourceName() string {
	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", c.Driver, c.Username, c.Password, c.Host, c.DB, c.SslMode)
}

func (c Database) GetConnectionMaxLifetime() time.Duration {
	return time.Duration(c.ConnectionMaxLifetime) * time.Minute
}

func (c Database) GetConnectionMaxIdleTime() time.Duration {
	return time.Duration(c.ConnectionMaxIdleTime) * time.Minute
}
