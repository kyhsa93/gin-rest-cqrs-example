package controller

import (
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

func Create(data *dto.Account, repository repository.Repository) {
	repository.Save(data)
}
