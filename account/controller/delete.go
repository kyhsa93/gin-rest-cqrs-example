package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
)

// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200
// @Router /accounts/{id} [delete]
// @Security AccessToken
// @Security RefreshToken
func (controller *Controller) delete(context *gin.Context) {
	controller.AuthenticateHTTPRequest(context)

	id := context.Param("id")

	command := &command.DeleteCommand{
		AccountID: id,
	}

	if err := controller.commandBus.Handle(command); err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, "Account deleted")
}
