package controller

import (
	"github.com/kyhsa93/go-rest-example/study/model"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func ReadItem(id string, repository repository.Repository) model.Study {
	var model model.Study
	result := repository.FindById(id)
	model.ID = result.ID
	model.Name = result.Name
	model.Description = result.Description
	model.CreatedAt = result.CreatedAt
	model.UpdatedAt = result.UpdatedAt
	return model
}

func ReadList(repository repository.Repository) model.Studies {
	var studies model.Studies
	var study model.Study

	result := repository.FindAll()
	for _, data := range result {
		study.ID = data.ID
		study.Name = data.Name
		study.Description = data.Description
		study.CreatedAt = data.CreatedAt
		study.UpdatedAt = data.UpdatedAt
		studies = append(studies, study)
	}
	return studies
}
