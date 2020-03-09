package config

import (
	"os"
)

// RedisConfiginterface redis config interface
type RedisConfiginterface interface {
	Address() string
	Password() string
}

// Redis redis struct
type Redis struct {
	address  string `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
	password string `env:"REDIS_PASSWORD" envDefault:""`
}

// NewRedisConfig create redis instance
func NewRedisConfig() *Redis {
	address := "localhost:6379"
	password := ""

	if env := os.Getenv("REDIS_ADDRESS"); env != "" {
		address = env
	}
	if env := os.Getenv("REDIS_PASSWORD"); env != "" {
		password = env
	}

	redis := &Redis{
		address:  address,
		password: password,
	}
	return redis
}

// Address get redis address
func (redis *Redis) Address() string {
	return redis.address
}

// Password get redis password
func (redis *Redis) Password() string {
	return redis.password
}
