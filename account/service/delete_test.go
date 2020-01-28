package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestDelete(t *testing.T) {
	repository := &mockedRepository{}
	config := &mockedConfig{}
	serviceInstance := service.New(repository, config)
	serviceInstance.Delete("accountID")
}
