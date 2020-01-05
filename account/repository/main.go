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
	if accountID == "" {
		_, count := repository.findByEmail(data.Email)
		if *count != 0 {
			return
		}
	}

	account := &entity.Account{}
	account.Email = data.Email
	account.SocialID = data.SocialID

	if accountID != "" {
		account.ID = accountID
	}

	err := repository.database.Save(account).Error

	if err != nil {
		panic(err)
	}
}

// FindByEmailAndSocialID find all account
func (repository *Repository) FindByEmailAndSocialID(email string, SocialID string) (data *model.Account) {
	var accountEntity entity.Account
	var accountModel model.Account

	err := repository.database.Where(&entity.Account{Email: email, SocialID: SocialID}).Take(&accountEntity).Error
	if err != nil {
		panic(err)
	}

	accountModel.ID = accountEntity.ID
	accountModel.Email = accountEntity.Email
	accountModel.CreatedAt = accountEntity.CreatedAt
	accountModel.UpdatedAt = accountEntity.UpdatedAt

	return &accountModel
}

// FindByID find account by accountId
func (repository *Repository) FindByID(id string) (data *model.Account) {
	var accountEntity entity.Account
	var accountModel model.Account

	repository.database.Where(&entity.Account{Model: entity.Model{ID: id}}).Take(&accountEntity)
	accountModel.ID = accountEntity.ID
	accountModel.CreatedAt = accountEntity.CreatedAt
	accountModel.UpdatedAt = accountEntity.UpdatedAt
	return &accountModel
}

// Delete delete account by accountId
func (repository *Repository) Delete(id string) {
	err := repository.database.Delete(&entity.Account{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}

func (repository *Repository) findByEmail(email string) (data *entity.Account, count *int) {
	var account entity.Account
	var accountCount int
	repository.database.Where(&entity.Account{Email: email}).Find(&account).Count(&accountCount)
	return &account, &accountCount
}
