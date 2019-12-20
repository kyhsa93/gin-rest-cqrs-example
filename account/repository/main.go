package repository

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/entity"
	"github.com/kyhsa93/go-rest-example/config"

	"github.com/jinzhu/gorm"
)

func database() *gorm.DB {
	return config.GetConnection()
}

// Save create or update account
func Save(data *dto.Account) {
	account := &entity.Account{}
	err := database().Save(account).Error
	if err != nil {
		panic(err)
	}
}

// FindAll find all account
func FindAll() entity.Accounts {
	var accounts entity.Accounts
	err := database().Find(&accounts).Error
	if err != nil {
		panic(err)
	}
	return accounts
}

// FindByID find account by accountId
func FindByID(id string) entity.Account {
	var account entity.Account
	database().Where(&entity.Account{Model: entity.Model{ID: id}}).Take(&account)
	return account
}

// Delete delete account by accountId
func Delete(id string) {
	err := database().Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
