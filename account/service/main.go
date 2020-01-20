package service

import (
	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/account/model"
	"github.com/kyhsa93/gin-rest-example/account/repository"
)

// Interface service interface
type Interface interface {
	Create(data *dto.Account)
	ReadAccountByID(acountID string) *model.Account
	ReadAccountByEmailAndSocialID(email string, socialID string) *model.Account
	Update(accountID string, data *dto.Account)
	Delete(accountID string)
}

// Service account service struct
type Service struct {
	repository repository.Interface
}

func (service *Service) entityToModel(entity entity.Account) *model.Account {
	var accountModel model.Account
	accountModel.ID = entity.ID
	accountModel.Email = entity.Email
	accountModel.CreatedAt = entity.CreatedAt
	accountModel.UpdatedAt = entity.UpdatedAt
	return &accountModel
}

// NewService create account service instance
func NewService(repository repository.Interface) *Service {
	return &Service{repository: repository}
}
