package controller

import (
	"go-rest-example/study/dto"
	"go-rest-example/study/repository"
)

func Update(id string, data *dto.Command, repository repository.Repository) {
	oldData := ReadItem(id, repository)
	if oldData.ID != id {
		return
	}
	repository.Save(&dto.Command{Name: data.Name, Description: data.Description})
}
