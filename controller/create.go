package controller

import (
	"go-rest-example/dto"
	"go-rest-example/repository"
)

func Create(data *dto.Study, repository repository.Repository) {
	repository.Save(data)
}
