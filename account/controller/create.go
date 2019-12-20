package controller

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

// Create create account
func Create(data *dto.Account) {
	repository.Save(data)
}
