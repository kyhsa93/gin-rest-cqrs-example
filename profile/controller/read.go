package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/query"
)

// @Tags Profiles
// @Accept json
// @Produce json
// @Param id path string true "profile id"
// @Success 200 {object} model.Profile
// @Router /profiles/{id} [get]
func (controller *Controller) readByID(context *gin.Context) {
	id := context.Param("id")
	query := &query.ReadProfileByIDQuery{ProfileID: id}
	profile, _ := controller.queryBus.Handle(query)

	if profile == nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, profile)
}
