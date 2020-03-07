package file

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

func getMongoDBClient(config *config.Config) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(
		"mongodb://" +
			config.Database.User + ":" + config.Database.Password +
			"@" + config.Database.Host + ":" + config.Database.Port,
	)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	client.Ping(context.TODO(), nil)
	collection := client.Database(
		config.Database.Name,
	).Collection("files")

	return collection
}

func getRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
	})
}

// InitializeFile init file module
func InitializeFile(engine *gin.Engine, config *config.Config, util *util.Util) {
	mongoClient := getMongoDBClient(config)
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, mongoClient)
	api := api.New(config)
	aws := aws.New(config)
	commandBus := command.New(repository, api, aws, config)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, api)
}
