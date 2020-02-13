package config

import (
	"github.com/caarlos0/env"
)

// Redis redis struct
type Redis struct {
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
}

// NewRedis create redis instance
func NewRedis() *Redis {
	redis := &Redis{}
	env.Parse(redis)
	return redis
}
