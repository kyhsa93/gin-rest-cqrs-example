package repository

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"

	"github.com/jinzhu/gorm"
)

// Interface repository inteface
type Interface interface {
	Save(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
		imageKey string,
		gender string,
		Interest string,
	)
	FindByEmailAndProvider(email string, provider string, unscoped bool) entity.Account
	FindByID(id string) entity.Account
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

func (repository *Repository) setCache(key string, accountEntity *entity.Account) {
	marshaledEntity, _ := json.Marshal(&accountEntity)
	setRedisDataError := repository.redis.Set(
		"account:"+key, string(marshaledEntity), time.Second,
	).Err()
	if setRedisDataError != nil {
		log.Println(setRedisDataError)
	}
}

func (repository *Repository) getCache(key string) *entity.Account {
	data, getDataFromRedisErrorByKey := repository.redis.Get("account:" + key).Result()
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

// Save create or update account
func (repository *Repository) Save(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
	imageKey string,
	gender string,
	interest string,
) {
	accountEntity := &entity.Account{
		Model:    entity.Model{ID: accountID},
		Email:    email,
		Provider: provider,
		Password: password,
		SocialID: socialID,
		ImageKey: imageKey,
		Gender:   gender,
		Interest: interest,
	}

	err := repository.database.Save(accountEntity).Error

	if err != nil {
		panic(err)
	}
	repository.setCache(accountID, accountEntity)
}

// FindByEmailAndProvider find all account
func (repository *Repository) FindByEmailAndProvider(
	email string,
	provider string,
	unscoped bool,
) entity.Account {
	accountEntity := entity.Account{}
	condition := entity.Account{Email: email, Provider: provider}

	if unscoped == true {
		repository.database.Unscoped().Where(&condition).First(&accountEntity)
		return accountEntity
	}

	if cache := repository.getCache(email); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&accountEntity)
	repository.setCache(email, &accountEntity)

	return accountEntity
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) entity.Account {
	accountEntity := entity.Account{}
	condition := entity.Account{Model: entity.Model{ID: id}}

	if cache := repository.getCache(id); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&accountEntity)
	repository.setCache(id, &accountEntity)
	return accountEntity
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	condition := &entity.Account{Model: entity.Model{ID: id}}
	err := repository.database.Delete(condition).Error
	if err != nil {
		panic(err)
	}
}
