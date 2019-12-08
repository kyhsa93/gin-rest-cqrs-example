package controller

import (
	"study/model"
	"study/repository"
)

func Create(study *model.Study, repository repository.Repository) {
	repository.Save(study)
}
