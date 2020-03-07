package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/dto"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/query"
)

// @Description update profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param UpdateProfile body body.UpdateProfile true "update profile data"
// @Success 200 {object} model.Profile
// @Router /profiles [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.UpdateProfile
	data.AccountID = account.ID

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadProfileByAccountIDQuery{
		AccountID: data.AccountID,
	}
	profile, err := controller.queryBus.Handle(query)
	if profile.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	command := &command.UpdateProfileCommand{
		ID:                    profile.ID,
		FileID:                data.FileID,
		InterestedField:       data.InterestedField,
		InterestedFieldDetail: data.InterestedFieldDetail,
	}
	updatedProfile, err := controller.commandBus.Handle(command)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, updatedProfile)
}
