package router

import (
	"go-rest-example/study/controller"
	"go-rest-example/study/dto"
	"go-rest-example/study/repository"

	"github.com/gin-gonic/gin"
)

// @Description create study group
// @Tags Studies
// @Accept  json
// @Produce  json
// @Param study body dto.Command true "Add study"
// @Success 200
// @Router /studies [post]
func Create(context *gin.Context, repository repository.Repository) {
	var data dto.Command
	context.ShouldBindJSON(&data)
	controller.Create(&data, repository)
}
