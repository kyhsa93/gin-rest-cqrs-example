package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/entity"
	"github.com/kyhsa93/gin-rest-example/account/service"
	"github.com/kyhsa93/gin-rest-example/config/auth"
	"github.com/kyhsa93/gin-rest-example/config/database"
	"github.com/kyhsa93/gin-rest-example/config/email"
	"github.com/kyhsa93/gin-rest-example/config/redis"
	"github.com/kyhsa93/gin-rest-example/config/server"
)

type mockedRepository struct{}

func (repository *mockedRepository) Save(accountID string, email string, provider string, socialID string, password string) {
}
func (repository *mockedRepository) FindByEmailAndProvider(email string, provider string, unscoped bool) entity.Account {
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

type mockedConfig struct{}

func (config *mockedConfig) Auth() *auth.Auth             { return &auth.Auth{} }
func (config *mockedConfig) Server() *server.Server       { return &server.Server{} }
func (config *mockedConfig) Database() *database.Database { return &database.Database{} }
func (config *mockedConfig) Redis() *redis.Redis          { return &redis.Redis{} }
func (config *mockedConfig) Email() *email.Email          { return &email.Email{} }

func TestNew(t *testing.T) {
	repository := &mockedRepository{}
	config := &mockedConfig{}
	serviceInstance := service.New(repository, config)
	if serviceInstance == nil {
		t.Error("Can not create service instance")
	}
}
