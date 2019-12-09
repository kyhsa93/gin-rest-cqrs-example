package controller

import (
	"study/dto"
	"study/repository"
)

func Create(data *dto.Study, repository repository.Repository) {
	repository.Save(data)
}
