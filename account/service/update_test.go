package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestUpdate(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.NewService(repository)
	data := &dto.Account{}
	serviceInstance.Update("accountID", data)
}
