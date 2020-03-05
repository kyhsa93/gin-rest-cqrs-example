package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/dto"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/query"
)

// @Description create profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profile body body.CreateProfile true "Create Profile data"
// @Success 201 {object} model.Profile
// @Router /profiles [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.CreateProfile

	data.AccountID = account.ID

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if data.Email == "" || data.Gender == "" || data.InterestedField == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included.")
		return
	}

	query := &query.ReadProfileByAccountIDQuery{
		AccountID: data.AccountID,
	}
	alreadyExisted, err := controller.queryBus.Handle(query)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	if alreadyExisted.ID != "" {
		httpError := controller.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), "Profile is already existed.")
		return
	}

	command := &command.CreateCommand{
		Email:                 data.Email,
		AccountID:             data.AccountID,
		Gender:                data.Gender,
		InterestedField:       data.InterestedField,
		InterestedFieldDetail: data.InterestedFieldDetail,
	}

	createdProfile, handlingError := controller.commandBus.Handle(command)
	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, createdProfile)
}
