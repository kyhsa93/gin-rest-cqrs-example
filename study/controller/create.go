package controller

import (
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func Create(data *dto.Command, repository repository.Repository) {
	repository.Save(data)
}
