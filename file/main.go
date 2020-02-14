package file

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/repository"
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
	connection.AutoMigrate(&entity.File{})
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

// InitializeFile init file module
func InitializeFile(engine *gin.Engine, config *config.Config, util *util.Util) {
	databaseConnection := getDatabaseConnection(config)
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, databaseConnection)
	aws := aws.New(config)
	commandBus := command.New(repository, aws)
	controller.New(engine, commandBus, util)
}
