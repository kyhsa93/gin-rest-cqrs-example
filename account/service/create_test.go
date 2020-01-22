package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestCreate(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.New(repository)
	data := &dto.Account{}
	serviceInstance.Create(data.Email, data.Provider, data.SocialID, data.Password)
}
