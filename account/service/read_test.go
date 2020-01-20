package service_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-example/account/service"
)

func TestReadAccountByID(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.New(repository)
	result := serviceInstance.ReadAccountByID("accountID")
	if result != nil {
		t.Error("Read account by ID service is error")
	}
}

func TestReadAccountByEmailAndSocialID(t *testing.T) {
	repository := &mockedRepository{}
	serviceInstance := service.New(repository)
	result := serviceInstance.ReadAccountByEmailAndSocialID("email", "socialID")
	if result != nil {
		t.Error("Read account by email and socialID")
	}
}
