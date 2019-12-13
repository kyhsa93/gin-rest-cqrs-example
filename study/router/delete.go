package router

import (
	"github.com/gin-gonic/gin"
	"go-rest-example/study/controller"
	"go-rest-example/study/repository"
)

// @Description delete study by id
// @Tags Studies
// @Param id path string true "Study ID"
// @Success 200
// @Router /studies/{id} [delete]
func Delete(context *gin.Context, repository repository.Repository) {
	controller.Delete(context.Param("id"), repository)
}
