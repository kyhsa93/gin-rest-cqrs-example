package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/dto"
)

// @Description create file
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param usage formData string true "file usage"
// @Param image formData file false "Profile image file"
// @Success 201
// @Router /files [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	usage := context.PostForm("usage")
	image, _ := context.FormFile("image")
	dto := dto.File{
		AccountID: account.ID,
		Usage:     usage,
		File:      image,
	}

	if !dto.ValidateUsage() || image == nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.CreateCommand{
		AccountID:   account.ID,
		AccessToken: accessToken,
		Usage:       usage,
		Image:       image,
	}
	createdFile, handlingError := controller.commandBus.Handle(command)
	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, createdFile)
}
