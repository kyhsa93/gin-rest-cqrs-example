package controller

import (
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func ReadItem(id string, repository repository.Repository) dto.Study {
	var dto dto.Study
	result := repository.FindById(id)
	dto.ID = result.ID
	dto.Name = result.Name
	dto.Description = result.Description
	dto.CreatedAt = result.CreatedAt
	dto.UpdatedAt = result.UpdatedAt
	return dto
}

func ReadList(repository repository.Repository) dto.Studies {
	var studies dto.Studies
	var study dto.Study

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
