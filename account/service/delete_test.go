package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestDelete(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.New(repository)
	serviceInstance.Delete("accountID")
}
