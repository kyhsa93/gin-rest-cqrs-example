package service

import (
	"github.com/kyhsa93/go-rest-example/account/model"
)

// ReadAccount read account by acountID
func (service *Service) ReadAccount(acountID string) (data *model.Account) {
	account := service.repository.FindByID(acountID)

	if account.ID == "" {
		return nil
	}

	return account
}

// ReadAccountByEmailAndSocialID read account list
func (service *Service) ReadAccountByEmailAndSocialID(email string, socialID string) (data *model.Account) {
	account := service.repository.FindByEmailAndSocialID(email, socialID)

	if account.ID == "" {
		return nil
	}

	return account
}
