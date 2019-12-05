package config

import "github.com/caarlos0/env"

type Service struct {
	Port string `env:"PORT" envDefault:"5000"`
}

func GetService() *Service {
	service := Service{}
	env.Parse(&service)
	return &service
}
