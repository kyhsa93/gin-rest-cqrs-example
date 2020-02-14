package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
)

// @Description create file
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param accountID formData string true "accountId"
// @Param usage formData string true "file usage"
// @Param image formData file false "Profile image file"
// @Success 201
// @Router /files [post]
func (controller *Controller) create(context *gin.Context) {
	accountID := context.PostForm("account_id")
	usage := context.PostForm("usage")
	image, _ := context.FormFile("image")

	command := &command.CreateCommand{
		AccountID: accountID,
		Usage:     usage,
		Image:     image,
	}

	if err := controller.commandBus.Handle(command); err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
}
