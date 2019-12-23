package service

import (
	"github.com/kyhsa93/go-rest-example/account/model"
)

// ReadAccount read account by acountID
func (service *Service) ReadAccount(acountID string) model.Account {
	var model model.Account
	result := service.repository.FindByID(acountID)
	model.ID = result.ID
	model.CreatedAt = result.CreatedAt
	model.UpdatedAt = result.UpdatedAt
	return model
}

// ReadAccounts read account list
func (service *Service) ReadAccounts() model.Accounts {
	var accounts model.Accounts
	var account model.Account
	result := service.repository.FindAll()

	for _, data := range result {
		account.ID = data.ID
		account.CreatedAt = data.CreatedAt
		account.UpdatedAt = data.UpdatedAt
		accounts = append(accounts, account)
	}
	return accounts
}
