package profile

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoDBClient(config *config.Config) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(
		"mongodb://" +
			config.Database.User() + ":" + config.Database.Password() +
			"@" + config.Database.Host() + ":" + config.Database.Port(),
	)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	client.Ping(context.TODO(), nil)
	collection := client.Database(
		config.Database.Name(),
	).Collection("profiles")

	return collection
}

func getRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
	})
}

// Initialize initialize profile module
func Initialize(
	engine *gin.Engine, config *config.Config, util *util.Util,
) {
	mongoClient := getMongoDBClient(config)
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, mongoClient)
	api := api.New(config)
	commandBus := command.New(repository, config)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, config, api)
}
