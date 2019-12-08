package controller

import (
	"study/model"
	"study/repositories"
)

func Create(study *model.Study, repository repositories.Repository) {
	repository.Save(study)
}
