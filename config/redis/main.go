package redis

import (
	"encoding/json"
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/entity"
)

type redisEnvironmentValue struct {
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
}

// Interface interface for redis client
type Interface interface {
	Set(key string, accountEntity *entity.Account)
	Get(key string) *entity.Account
}

// Redis redis struct
type Redis struct {
	client *redis.Client
}

// Set set data in redis
func (cache *Redis) Set(key string, accountEntity *entity.Account) {
	marshaledEntity, _ := json.Marshal(&accountEntity)
	setRedisDataError := cache.client.Set("account:"+key, string(marshaledEntity), time.Second).Err()
	if setRedisDataError != nil {
		log.Println(setRedisDataError)
	}
}

// Get get data from redis by accountID
func (cache *Redis) Get(key string) *entity.Account {
	data, getDataFromRedisErrorByKey := cache.client.Get("account:" + key).Result()
	if getDataFromRedisErrorByKey != nil {
		log.Println(getDataFromRedisErrorByKey)
		return nil
	}

	entity := &entity.Account{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if jsonUnmarshalError != nil {
		log.Println(jsonUnmarshalError)
		return nil
	}

	if entity.ID == "" {
		return nil
	}
	return entity
}

func getRedisConfig() *redis.Options {
	config := &redisEnvironmentValue{}
	env.Parse(config)
	return &redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
	}
}

func getClient() *redis.Client {
	return redis.NewClient(getRedisConfig())
}

// New create redis instance
func New() *Redis {
	return &Redis{client: getClient()}
}
