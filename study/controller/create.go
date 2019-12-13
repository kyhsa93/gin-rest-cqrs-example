package controller

import (
	"go-rest-example/study/dto"
	"go-rest-example/study/repository"
)

func Create(data *dto.Study, repository repository.Repository) {
	repository.Save(data)
}
