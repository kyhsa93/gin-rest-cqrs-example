package controller

import "go-rest-example/repository"

func Delete(id string, repository repository.Repository) {
	repository.Delete(id)
}
