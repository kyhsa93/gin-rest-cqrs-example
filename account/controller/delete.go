package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
)

// @Description delete account
// @Tags Account
// @Success 200 {object} model.Account
// @Router /accounts [delete]
// @Security AccessToken
func (controller *Controller) delete(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.DeleteCommand{
		AccountID: account.ID,
	}

	deletedAccount, handlingError := controller.commandBus.Handle(command)

	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, deletedAccount)
}
