package service

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
)

// Update update account by accountID
func (service *Service) Update(accountID string, data *dto.Account) {
	oldData := service.ReadAccountByID(accountID)
	if oldData.ID != accountID {
		return
	}
	service.repository.Save(data, accountID)
}
