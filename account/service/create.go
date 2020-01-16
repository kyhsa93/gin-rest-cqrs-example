package service

import (
	"github.com/kyhsa93/gin-rest-example/account/dto"
)

// Create create account
func (service *Service) Create(data *dto.Account) {
	service.repository.Save(data, "")
}
