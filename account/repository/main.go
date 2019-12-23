package repository

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/entity"
	"github.com/kyhsa93/go-rest-example/account/model"
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
	database.AutoMigrate(&entity.Account{})
	return &Repository{database: database}
}

// Save create or update account
func (repository *Repository) Save(data *dto.Account, accountID string) {
	account := &entity.Account{}
	account.Email = data.Email
	account.Password = data.Password
	if accountID != "" {
		account.ID = accountID
	}
	err := repository.database.Save(account).Error
	if err != nil {
		panic(err)
	}
}

// FindAll find all account
func (repository *Repository) FindAll(email string, password string) model.Accounts {
	var accountEntities entity.Accounts
	var accountModel model.Account
	var accountModels model.Accounts

	err := repository.database.Where(&entity.Account{Email: email, Password: password}).Find(&accountEntities).Error
	if err != nil {
		panic(err)
	}

	for _, accountEntity := range accountEntities {
		accountModel.ID = accountEntity.ID
		accountModel.Email = accountEntity.Email
		accountModel.CreatedAt = accountEntity.CreatedAt
		accountModel.UpdatedAt = accountEntity.UpdatedAt
		accountModels = append(accountModels, accountModel)
	}
	return accountModels
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) model.Account {
	var accountEntity entity.Account
	var accountModel model.Account
	repository.database.Where(&entity.Account{Model: entity.Model{ID: id}}).Take(&accountEntity)
	accountModel.ID = accountEntity.ID
	accountModel.CreatedAt = accountEntity.CreatedAt
	accountModel.UpdatedAt = accountEntity.UpdatedAt
	return accountModel
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	err := repository.database.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
