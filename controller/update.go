package controller

import (
	"go-rest-example/dto"
	"go-rest-example/repository"
)

func Update(id string, data *dto.Study, repository repository.Repository) {
	oldData := ReadItem(id, repository)
	if oldData.ID != id {
		return
	}
	repository.Save(&dto.Study{Name: data.Name, Description: data.Description})
}
