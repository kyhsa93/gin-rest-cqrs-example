package service

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/model"
)

// Create create account
func (service *Service) Create(data *dto.Account) *model.Account {
	account := service.repository.FindByEmail(data.Email)

	if account != nil {
		return account
	}

	service.repository.Save(data, "")

	return nil
}
