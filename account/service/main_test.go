package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/account/service"
)

type mockedRepository struct{}

func (repository *mockedRepository) Save(accountID string, email string, provider string, socialID string, password string) {
}
func (repository *mockedRepository) FindByEmailAndSocialID(email string, provider string, socialID string, password string) entity.Account {
	return entity.Account{}
}
func (repository *mockedRepository) FindByID(id string) entity.Account {
	return entity.Account{}
}
func (repository *mockedRepository) Update(accountID string, data *dto.Account) {}
func (repository *mockedRepository) Delete(id string)                           {}
func (repository *mockedRepository) dtoToEntity(dto *dto.Account) *entity.Account {
	return &entity.Account{}
}

func TestNew(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.New(repository)
	if serviceInstance == nil {
		t.Error("Can not create service instance")
	}
}
