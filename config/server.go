package config

import "fmt"

type Server struct {
	Port    string `mapstructure:"port"`
	BaseUrl string `mapstructure:"common-url"`
}

func (c Server) Address() string {
	return fmt.Sprintf("%s:%s", c.BaseUrl, c.Port)
}
