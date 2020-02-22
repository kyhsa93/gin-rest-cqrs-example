package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/query"
)

// @Tags Files
// @Accept json
// @Produce json
// @Param id path string true "file Id"
// @Success 200 {object} model.File
// @Router /files/{id} [get]
// @Security AccessToken
func (controller *Controller) readFileByID(context *gin.Context) {
	id := context.Param("id")
	query := &query.ReadFileByIDQuery{FileID: id}
	file, _ := controller.queryBus.Handle(query)

	if file == nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, file)
}
