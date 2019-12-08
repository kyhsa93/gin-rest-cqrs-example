package controller

import (
	"study/model"
	"study/repositories"
)

func ReadItem(id string, repository repositories.Repository) model.Study {
	return repository.FindById(id)
}

func ReadList(repository repositories.Repository) model.Studies {
	return repository.FindAll()
}
