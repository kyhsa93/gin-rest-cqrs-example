package controller

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

// Update update account by accountID
func Update(accountID string, data *dto.Account) {
	oldData := ReadItem(accountID)
	if oldData.ID != accountID {
		return
	}
	repository.Save(&dto.Account{})
}
