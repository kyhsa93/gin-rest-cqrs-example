package controller

import "study/model"

import "study/repositories"

func Update(id string, study *model.Study, repository repositories.Repository) {
	oldData := ReadItem(id, repository)
	if oldData.ID != id {
		return
	}
	oldData.Name = study.Name
	oldData.Description = study.Description
	repository.Save(&oldData)
}
