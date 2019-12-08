package controller

import "study/repositories"

func Delete(id string, repository repositories.Repository) {
	repository.Delete(id)
}
