package account

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/gin-rest-example/account/application/command"
	"github.com/kyhsa93/gin-rest-example/account/application/query"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/entity"
	"github.com/kyhsa93/gin-rest-example/account/interface/controller"
	"github.com/kyhsa93/gin-rest-example/config"
	"github.com/kyhsa93/gin-rest-example/util"
)

func getDatabaseConnection(config *config.Config) *gorm.DB {
	user := config.Database.User
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	name := config.Database.Name
	logging := config.Database.Logging

	connection, err := gorm.Open(
		"mysql", user+":"+password+"@tcp("+host+":"+port+")/"+name+"?parseTime=true",
	)
	if err != nil {
		panic(err)
	}
	connection.LogMode(logging)
	connection.AutoMigrate(&entity.Account{})
	return connection
}

func getRedisClient(config *config.Config) *redis.Client {
	host := config.Redis.Host
	port := config.Redis.Password
	password := config.Redis.Password
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
	})
}

// InitializeAccount innitialize account module
func InitializeAccount(engine *gin.Engine, config *config.Config, util *util.Util) {
	databaseConnection := getDatabaseConnection(config)
	redisClient := getRedisClient(config)
	infra := infrastructure.New(databaseConnection, redisClient, config)
	commandBus := command.New(infra)
	queryBus := query.New(infra, config)
	controller.New(engine, commandBus, queryBus, util)
}
