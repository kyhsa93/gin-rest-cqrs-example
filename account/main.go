package account

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/email"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoDBClient() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://root:test@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	client.Ping(context.TODO(), nil)
	collection := client.Database("gin-rest-cqrs-example").Collection("accounts")

	return collection
}

func getRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
	})
}

// InitializeAccount innitialize account module
func InitializeAccount(
	engine *gin.Engine, config *config.Config, util *util.Util,
) {
	mongoClient := getMongoDBClient()
	redisClient := getRedisClient(config)
	repository := repository.New(redisClient, mongoClient)
	email := email.New(config)
	aws := aws.New(config)
	api := api.New(config)
	commandBus := command.New(repository, email, aws, config, api)
	queryBus := query.New(config, repository)
	controller.New(engine, commandBus, queryBus, util, config, api)
}
