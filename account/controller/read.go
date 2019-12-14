package controller

import (
	"github.com/kyhsa93/go-rest-example/account/model"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

func ReadItem(id string, repository repository.Repository) model.Account {
	var model model.Account
	result := repository.FindById(id)
	model.ID = result.ID
	model.CreatedAt = result.CreatedAt
	model.UpdatedAt = result.UpdatedAt
	return model
}

func ReadList(repository repository.Repository) model.Accounts {
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
