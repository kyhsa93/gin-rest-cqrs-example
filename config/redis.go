package config

import (
	"github.com/caarlos0/env"
)

// Redis redis struct
type Redis struct {
	Address  string `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
}

// NewRedis create redis instance
func NewRedis() *Redis {
	redis := &Redis{}
	env.Parse(redis)
	return redis
}
