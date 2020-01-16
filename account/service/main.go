package service

import (
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/account/model"
	"github.com/kyhsa93/gin-rest-example/account/repository"
)

// Service account service struct
type Service struct {
	repository *repository.Repository
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
func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
