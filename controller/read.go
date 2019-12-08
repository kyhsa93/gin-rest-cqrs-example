package controller

import (
	"study/model"
	"study/repositories"

	"github.com/gin-gonic/gin"
)

func ReadItem(context *gin.Context, repository repositories.Repository) model.Study {
	return repository.FindById(context.Param("id"))
}

func ReadList(repository repositories.Repository) model.Studies {
	return repository.FindAll()
}
