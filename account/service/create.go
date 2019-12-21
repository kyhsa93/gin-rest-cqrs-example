package service

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
)

// Create create account
func (service *Service) Create(data *dto.Account) {
	service.repository.Save(data)
}
