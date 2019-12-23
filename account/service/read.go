package service

import (
	"github.com/kyhsa93/go-rest-example/account/model"
)

// ReadAccount read account by acountID
func (service *Service) ReadAccount(acountID string) model.Account {
	return service.repository.FindByID(acountID)
}

// ReadAccounts read account list
func (service *Service) ReadAccounts(email string, password string) model.Accounts {
	return service.repository.FindAll(email, password)
}
