package repository

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/entity"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) Save(data *dto.Account) {
	account := &entity.Account{}
	err := repository.db.Save(account).Error
	if err != nil {
		panic(err)
	}
}

func (repository *Repository) FindAll() entity.Accounts {
	var accounts entity.Accounts
	err := repository.db.Find(&accounts).Error
	if err != nil {
		panic(err)
	}
	return accounts
}

func (repository *Repository) FindById(id string) entity.Account {
	var account entity.Account
	repository.db.Where(&entity.Account{Model: entity.Model{ID: id}}).Take(&account)
	return account
}

func (repository *Repository) Delete(id string) {
	err := repository.db.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
