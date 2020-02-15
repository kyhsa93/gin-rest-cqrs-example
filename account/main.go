package account

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/email"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
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
	repository := repository.New(redisClient, databaseConnection)
	email := email.New(config)
	aws := aws.New(config)
	commandBus := command.New(repository, email, aws, config)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util)
}
