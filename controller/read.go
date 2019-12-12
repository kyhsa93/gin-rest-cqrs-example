package controller

import (
	"go-rest-example/model"
	"go-rest-example/repository"
)

func ReadItem(id string, repository repository.Repository) model.Study {
	return repository.FindById(id)
}

func ReadList(repository repository.Repository) model.Studies {
	return repository.FindAll()
}
