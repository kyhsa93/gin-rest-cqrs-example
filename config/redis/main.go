package redis

import (
	"github.com/caarlos0/env"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
)

// Interface interface for redis client
type Interface interface {
	Set(key string, accountEntity *entity.Account)
	Get(key string) *entity.Account
}

// Redis redis struct
type Redis struct {
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
}

// New create redis instance
func New() *Redis {
	redis := &Redis{}
	env.Parse(redis)
	return redis
}
