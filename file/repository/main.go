package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface repository inteface
type Interface interface {
	Create(
		fileID string,
		accountID string,
		usage string,
	) (entity.File, error)
	FindByID(fileID string) entity.File
}

// Repository repository for query to database
type Repository struct {
	redis *redis.Client
	mongo *mongo.Collection
}

// New create repository instance
func New(redis *redis.Client, mongo *mongo.Collection) Interface {
	return &Repository{mongo: mongo, redis: redis}
}

func (repository *Repository) setCache(key string, fileEntity *entity.File) {
	marshaledEntity, _ := json.Marshal(&fileEntity)
	setRedisDataError := repository.redis.Set(
		"file:"+key, string(marshaledEntity), time.Second,
	).Err()
	if setRedisDataError != nil {
		log.Println(setRedisDataError)
	}
}

func (repository *Repository) getCache(key string) *entity.File {
	data, getDataFromRedisErrorByKey := repository.redis.Get("file:" + key).Result()
	if getDataFromRedisErrorByKey != nil {
		log.Println(getDataFromRedisErrorByKey)
		return nil
	}

	entity := &entity.File{}
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

// Create create file
func (repository *Repository) Create(
	fileID string,
	accountID string,
	usage string,
) (entity.File, error) {
	fileEntity := entity.File{
		ID:        fileID,
		AccountID: accountID,
		Usage:     usage,
		CreatedAt: time.Now(),
	}

	insertResult, err := repository.mongo.InsertOne(
		context.TODO(),
		fileEntity,
	)
	if err != nil || insertResult == nil {
		return fileEntity, err
	}
	repository.setCache(accountID, &fileEntity)
	return fileEntity, nil
}

// FindByID fine file data using file id
func (repository *Repository) FindByID(fileID string) entity.File {
	if cache := repository.getCache(fileID); cache != nil {
		return *cache
	}
	fileEntity := entity.File{}
	repository.mongo.FindOne(
		context.TODO(),
		bson.M{"_id": fileID},
	).Decode(&fileEntity)
	repository.setCache(fileID, &fileEntity)
	return fileEntity
}
