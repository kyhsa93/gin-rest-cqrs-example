package controller

import "github.com/kyhsa93/go-rest-example/account/repository"

func Delete(id string, repository repository.Repository) {
	repository.Delete(id)
}
