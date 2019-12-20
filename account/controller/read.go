package controller

import (
	"github.com/kyhsa93/go-rest-example/account/model"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

// ReadItem read account by acountID
func ReadItem(acountID string) model.Account {
	var model model.Account
	result := repository.FindByID(acountID)
	model.ID = result.ID
	model.CreatedAt = result.CreatedAt
	model.UpdatedAt = result.UpdatedAt
	return model
}

// ReadList read account list
func ReadList() model.Accounts {
	var accounts model.Accounts
	var account model.Account

	result := repository.FindAll()

	for _, data := range result {
		account.ID = data.ID
		account.CreatedAt = data.CreatedAt
		account.UpdatedAt = data.UpdatedAt
		accounts = append(accounts, account)
	}
	return accounts
}
