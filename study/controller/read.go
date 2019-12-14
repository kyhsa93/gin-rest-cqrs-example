package controller

import (
	"github.com/kyhsa93/go-rest-example/study/model"
	"github.com/kyhsa93/go-rest-example/study/repository"
)

func ReadItem(id string, repository repository.Repository) model.Study {
	return repository.FindById(id)
}

func ReadList(repository repository.Repository) model.Studies {
	return repository.FindAll()
}
