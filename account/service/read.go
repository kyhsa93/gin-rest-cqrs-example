package service

import (
	"github.com/kyhsa93/gin-rest-example/account/model"
)

// ReadAccountByID read account by acountID
func (service *Service) ReadAccountByID(acountID string) *model.Account {
	entity := service.repository.FindByID(acountID)

	if entity.ID == "" {
		return nil
	}

	return service.entityToModel(entity)
}

// ReadAccountByEmailAndSocialID read account list
func (service *Service) ReadAccountByEmailAndSocialID(email string, provider string, socialID string, password string) *model.Account {
	entity := service.repository.FindByEmailAndSocialID(email, provider, socialID, password)

	if entity.ID == "" {
		return nil
	}

	return service.entityToModel(entity)
}
