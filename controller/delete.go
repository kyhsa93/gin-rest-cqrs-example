package controller

import "study/repository"

func Delete(id string, repository repository.Repository) {
	repository.Delete(id)
}
