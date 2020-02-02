package repository

import (
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/config"
	"github.com/kyhsa93/gin-rest-example/config/redis"

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
		Intereste string,
	)
	FindByEmailAndProvider(email string, provider string, unscoped bool) entity.Account
	FindByID(id string) entity.Account
	Delete(id string)
}

// Repository repository for query to database
type Repository struct {
	redis    redis.Interface
	database *gorm.DB
}

// New create repository instance
func New(config *config.Config) *Repository {
	database := config.Database.Connection
	database.AutoMigrate(&entity.Account{})
	redis := config.Redis
	return &Repository{database: database, redis: redis}
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
	intereste string,
) {
	accountEntity := &entity.Account{
		Model:     entity.Model{ID: accountID},
		Email:     email,
		Provider:  provider,
		Password:  password,
		SocialID:  socialID,
		ImageKey:  imageKey,
		Gender:    gender,
		Intereste: intereste,
	}

	err := repository.database.Save(accountEntity).Error

	if err != nil {
		panic(err)
	}
	repository.redis.Set(accountID, accountEntity)
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

	if cache := repository.redis.Get(email); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&accountEntity)
	repository.redis.Set(email, &accountEntity)

	return accountEntity
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) entity.Account {
	accountEntity := entity.Account{}
	condition := entity.Account{Model: entity.Model{ID: id}}

	if cache := repository.redis.Get(id); cache != nil {
		return *cache
	}
	repository.database.Where(&condition).First(&accountEntity)
	repository.redis.Set(id, &accountEntity)
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
