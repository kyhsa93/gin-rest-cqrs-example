package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
)

// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [delete]
// @Security AccessToken
func (controller *Controller) delete(context *gin.Context) {
	controller.AuthenticateHTTPRequest(context)

	id := context.Param("id")

	command := &command.DeleteCommand{
		AccountID: id,
	}

	deletedAccount, handlingError := controller.commandBus.Handle(command)

	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, deletedAccount)
}
