package repository

import (
	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/config"

	"github.com/jinzhu/gorm"
)

// Interface repository inteface
type Interface interface {
	Save(accountID string, email string, provider string, socialID string, password string)
	FindByEmailAndSocialID(email string, provider string, socialID string, password string, unscoped bool) entity.Account
	FindByID(id string) entity.Account
	Delete(id string)
}

// Repository repository for query to database
type Repository struct {
	database *gorm.DB
}

// New create repository instance
func New(config *config.Config) *Repository {
	database := config.Database.Connection
	database.AutoMigrate(&entity.Account{})
	return &Repository{database: database}
}

func (repository *Repository) dtoToEntity(data *dto.Account) *entity.Account {
	return &entity.Account{Email: data.Email, Provider: data.Provider, Password: data.Password, SocialID: data.SocialID}
}

// Save create or update account
func (repository *Repository) Save(accountID string, email string, provider string, socialID string, password string) {
	accountEntity := &entity.Account{
		Model:    entity.Model{ID: accountID},
		Email:    email,
		Provider: provider,
		Password: password,
		SocialID: socialID,
	}

	err := repository.database.Save(accountEntity).Error

	if err != nil {
		panic(err)
	}
}

// FindByEmailAndSocialID find all account
func (repository *Repository) FindByEmailAndSocialID(email string, provider string, socialID string, password string, unscoped bool) entity.Account {
	accountEntity := entity.Account{}
	condition := &entity.Account{Email: email, SocialID: socialID}

	if unscoped == true {
		repository.database.Unscoped().Where(condition).First(&accountEntity)
		return accountEntity
	}
	repository.database.Where(condition).First(&accountEntity)

	return accountEntity
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) entity.Account {
	accountEntity := entity.Account{}
	repository.database.Where(&entity.Account{Model: entity.Model{ID: id}}).First(&accountEntity)
	return accountEntity
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	err := repository.database.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
