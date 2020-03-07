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

// @Tags Profiles
// @Accept json
// @Produce json
// @Success 200 {object} model.Profile
// @Router /profiles [get]
// @Security AccessToken
func (controller *Controller) read(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	query := &query.ReadProfileByAccountIDQuery{
		AccountID: account.ID,
	}
	profile, err := controller.queryBus.Handle(query)
	if err != nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, profile)
}
