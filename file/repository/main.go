package repository

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/entity"

	"github.com/jinzhu/gorm"
)

// Interface repository inteface
type Interface interface {
	Save(
		fileID string,
		accountID string,
		usage string,
		imageKey string,
	)
	Delete(id string)
}

// Repository repository for query to database
type Repository struct {
	redis    *redis.Client
	database *gorm.DB
}

// New create repository instance
func New(redis *redis.Client, database *gorm.DB) *Repository {
	return &Repository{database: database, redis: redis}
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

// Save create or update account
func (repository *Repository) Save(
	fileID string,
	accountID string,
	usage string,
	imageKey string,
) {
	fileEntity := &entity.File{
		Model:     entity.Model{ID: fileID},
		AccountID: accountID,
		Usage:     usage,
		ImageKey:  imageKey,
	}

	err := repository.database.Save(fileEntity).Error

	if err != nil {
		panic(err)
	}
	repository.setCache(accountID, fileEntity)
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	condition := &entity.File{Model: entity.Model{ID: id}}
	err := repository.database.Delete(condition).Error
	if err != nil {
		panic(err)
	}
}
