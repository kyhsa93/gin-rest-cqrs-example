package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
)

// @Description create file
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param account_id formData string true "accountId"
// @Param usage formData string true "file usage"
// @Param image formData file false "Profile image file"
// @Success 201
// @Router /files [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	accountID := context.PostForm("account_id")
	auth := controller.AuthenticateHTTPReqeust(accessToken, accountID)
	if !auth {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	usage := context.PostForm("usage")
	image, _ := context.FormFile("image")

	if usage == "" || image == nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.CreateCommand{
		AccountID: accountID,
		Usage:     usage,
		Image:     image,
	}
	createdFile, handlingError := controller.commandBus.Handle(command)
	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, createdFile)
}
