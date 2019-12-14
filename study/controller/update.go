package controller

import (
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func Update(id string, data *dto.Study, repository repository.Repository) {
	oldData := ReadItem(id, repository)
	if oldData.ID != id {
		return
	}
	repository.Save(&dto.Study{Name: data.Name, Description: data.Description})
}
