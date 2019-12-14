package router

import (
	"github.com/kyhsa93/go-rest-example/study/controller"
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/repository"

	"github.com/gin-gonic/gin"
)

// @Description create study group
// @Tags Studies
// @Accept  json
// @Produce  json
// @param id path string true "Study ID"
// @Param study body dto.Study true "Add study"
// @Success 200
// @Router /studies/{id} [put]
func Update(context *gin.Context, repository repository.Repository) {
	var command dto.Study
	context.ShouldBindJSON(&command)
	controller.Update(context.Param("id"), &command, repository)
}
