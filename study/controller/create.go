package controller

import (
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func Create(data *dto.Study, repository repository.Repository) {
	repository.Save(data)
}
