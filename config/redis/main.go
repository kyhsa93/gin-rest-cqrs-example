package redis

import (
	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
)

type redisEnvironmentValue struct {
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
}

// Redis redis struct
type Redis struct {
	Client *redis.Client
}

func getRedisConfig() *redis.Options {
	config := &redisEnvironmentValue{}
	env.Parse(config)
	return &redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
	}
}

// NewClient create new redis client
func getClient() *redis.Client {
	return redis.NewClient(getRedisConfig())
}

// New create redis instance
func New() *Redis {
	return &Redis{Client: getClient()}
}
