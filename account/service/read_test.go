package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestReadAccountByID(t *testing.T) {
	repository := &mockedRepository{}
	config := &mockedConfig{}
	serviceInstance := service.New(repository, config)
	result := serviceInstance.ReadAccountByID("accountID")
	if result != nil {
		t.Error("Read account by ID service is error")
	}
}

func TestReadAccount(t *testing.T) {
	repository := &mockedRepository{}
	config := &mockedConfig{}
	serviceInstance := service.New(repository, config)
	result, err := serviceInstance.ReadAccount("email", "provider", "socialID", "password", false)
	if result != nil && err == nil {
		t.Error("Read account by email and socialID")
	}
}
