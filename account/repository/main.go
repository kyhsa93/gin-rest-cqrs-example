package repository

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"

	"github.com/jinzhu/gorm"
)

// Interface repository inteface
type Interface interface {
	TransactionStart() *gorm.DB
	TransactionCommit(transaction *gorm.DB)
	TransactionRollback(transaction *gorm.DB)
	Create(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
		imageKey string,
		gender string,
		Interest string,
		transaction *gorm.DB,
	) (entity.Account, error)
	Update(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
		imageKey string,
		gender string,
		Interest string,
		transaction *gorm.DB,
	) (entity.Account, error)
	FindByEmailAndProvider(
		email string, provider string, unscoped bool,
	) entity.Account
	FindByID(id string, unscoped bool) entity.Account
	Delete(id string, transaction *gorm.DB) entity.Account
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

func (repository *Repository) setCache(key string, accountEntity *entity.Account) {
	marshaledEntity, _ := json.Marshal(&accountEntity)
	setRedisDataError := repository.redis.Set(
		"account:"+key, string(marshaledEntity), time.Second,
	).Err()
	if setRedisDataError != nil {
		log.Println("Set Data to Redis Error: ", setRedisDataError)
	}
}

func (repository *Repository) getCache(key string) *entity.Account {
	data, getDataFromRedisError := repository.redis.Get("account:" + key).Result()
	if getDataFromRedisError != nil {
		log.Println("Get Data from Redis Error: ", getDataFromRedisError)
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

// Create create account
func (repository *Repository) Create(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
	imageKey string,
	gender string,
	interest string,
	transaction *gorm.DB,
) (entity.Account, error) {
	sameEmailAccount := entity.Account{}
	transaction.Where(entity.Account{Email: email}).First(sameEmailAccount)
	if sameEmailAccount.ID != "" {
		return sameEmailAccount, errors.New("Duplicated Email")
	}
	accountEntity := entity.Account{
		Model:    entity.Model{ID: accountID},
		Email:    email,
		Provider: provider,
		Password: password,
		SocialID: socialID,
		ImageKey: imageKey,
		Gender:   gender,
		Interest: interest,
	}
	insertError := transaction.Create(&accountEntity).Error
	if insertError != nil {
		repository.TransactionRollback(transaction)
		panic(insertError)
	}
	repository.setCache(accountID, &accountEntity)
	return accountEntity, nil
}

// Update update account
func (repository *Repository) Update(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
	imageKey string,
	gender string,
	interest string,
	transaction *gorm.DB,
) (entity.Account, error) {
	accountEntity := entity.Account{
		Model:    entity.Model{ID: accountID},
		Email:    email,
		Provider: provider,
		Password: password,
		SocialID: socialID,
		ImageKey: imageKey,
		Gender:   gender,
		Interest: interest,
	}

	err := transaction.Save(&accountEntity).Error
	if err != nil {
		repository.TransactionRollback(transaction)
		panic(err)
	}
	repository.setCache(accountID, &accountEntity)
	return accountEntity, nil
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
func (repository *Repository) FindByID(id string, unscoped bool) entity.Account {
	accountEntity := entity.Account{}
	condition := entity.Account{Model: entity.Model{ID: id}}

	if unscoped == true {
		repository.database.Unscoped().Where(&condition).First(&accountEntity)
		return accountEntity
	}

	if cache := repository.getCache(id); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&accountEntity)
	repository.setCache(id, &accountEntity)
	return accountEntity
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string, transaction *gorm.DB) entity.Account {
	accountEntity := entity.Account{}
	condition := &entity.Account{Model: entity.Model{ID: id}}
	err := transaction.Delete(condition).Error
	if err != nil {
		repository.TransactionRollback(transaction)
		panic(err)
	}
	transaction.Unscoped().Where(condition).First(&accountEntity)
	return accountEntity
}
