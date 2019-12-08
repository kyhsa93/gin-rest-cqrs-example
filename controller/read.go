package controller

import (
	"study/model"
	"study/repository"
)

func ReadItem(id string, repository repository.Repository) model.Study {
	return repository.FindById(id)
}

func ReadList(repository repository.Repository) model.Studies {
	return repository.FindAll()
}
