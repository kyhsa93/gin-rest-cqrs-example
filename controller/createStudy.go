package controller

import (
	"study/model"
	"study/repositories"
)

func CreateStudy(study *model.Study, repository repositories.StudyRepository) {

	repository.Save(study)
}
