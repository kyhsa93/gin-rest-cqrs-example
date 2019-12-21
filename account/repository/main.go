package repository

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/entity"
	"github.com/kyhsa93/go-rest-example/config"

	"github.com/jinzhu/gorm"
)

// Repository repository for query to database
type Repository struct {
	database *gorm.DB
}

// NewRepository create repository instance
func NewRepository(config *config.Config) *Repository {
	database := config.Database.Connection
	return &Repository{database: database}
}

// Save create or update account
func (repository *Repository) Save(data *dto.Account) {
	account := &entity.Account{}
	err := repository.database.Save(account).Error
	if err != nil {
		panic(err)
	}
}

// FindAll find all account
func (repository *Repository) FindAll() entity.Accounts {
	var accounts entity.Accounts
	err := repository.database.Find(&accounts).Error
	if err != nil {
		panic(err)
	}
	return accounts
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) entity.Account {
	var account entity.Account
	repository.database.Where(&entity.Account{Model: entity.Model{ID: id}}).Take(&account)
	return account
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	err := repository.database.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
