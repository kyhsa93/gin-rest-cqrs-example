package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/dto"
)

// @Description create profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param profile body dto.Profile true "Create Profile data"
// @Success 201 {object} model.Profile
// @Router /profiles [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	var data dto.Profile

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	accessToken := context.GetHeader("Authorization")
	auth := controller.AuthenticateHTTPReqeust(accessToken, data.AccountID)
	if !auth {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if data.Email == "" || data.Gender == "" || data.InterestedFeild == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included.")
		return
	}

	if !controller.ValidateFileID(data.AccountID, data.FileID) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Invalid fileId")
		return
	}

	command := &command.CreateCommand{
		Email:                 data.Email,
		AccountID:             data.AccountID,
		Gender:                data.Gender,
		FileID:                data.FileID,
		InterestedFeild:       data.InterestedFeild,
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
