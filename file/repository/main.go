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
	TransactionStart() *gorm.DB
	TransactionCommit(transaction *gorm.DB)
	TransactionRollback(transaction *gorm.DB)
	Create(
		fileID string,
		accountID string,
		usage string,
		transaction *gorm.DB,
	) (entity.File, error)
	FindByID(fileID string) entity.File
	Delete(id string, transaction *gorm.DB) entity.File
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

// TransactionStart start database transaction
func (repository *Repository) TransactionStart() *gorm.DB {
	return repository.database.Begin()
}

// TransactionCommit commit database transaction
func (repository *Repository) TransactionCommit(transaction *gorm.DB) {
	transaction.Commit()
}

// TransactionRollback rollback database transaction
func (repository *Repository) TransactionRollback(transaction *gorm.DB) {
	transaction.Rollback()
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
	transaction *gorm.DB,
) (entity.File, error) {
	fileEntity := entity.File{
		Model:     entity.Model{ID: fileID},
		AccountID: accountID,
		Usage:     usage,
	}

	insertError := transaction.Create(&fileEntity).Error
	if insertError != nil {
		repository.TransactionRollback(transaction)
		panic(insertError)
	}
	repository.setCache(accountID, &fileEntity)
	return fileEntity, nil
}

// FindByID fine file data using file id
func (repository *Repository) FindByID(fileID string) entity.File {
	fileEntity := entity.File{}
	condition := entity.File{Model: entity.Model{ID: fileID}}

	if cache := repository.getCache(fileID); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&fileEntity)
	repository.setCache(fileID, &fileEntity)
	return fileEntity
}

// Delete delete file by fileId
func (repository *Repository) Delete(id string, transaction *gorm.DB) entity.File {
	fileEntity := entity.File{}
	condition := &entity.File{Model: entity.Model{ID: id}}
	err := transaction.Delete(condition).Error
	if err != nil {
		repository.TransactionRollback(transaction)
		panic(err)
	}
	transaction.Unscoped().Where(condition).First(&fileEntity)
	return fileEntity
}
