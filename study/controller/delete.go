package controller

import "go-rest-example/study/repository"

func Delete(id string, repository repository.Repository) {
	repository.Delete(id)
}
