package controller

import (
	"study/model"
	"study/repositories"
)

func Create(study *model.Study, repository repositories.StudyRepository) {

	repository.Save(study)
}
