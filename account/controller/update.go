package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/dto"
)

// @Description update account
// @Tags Account
// @Accept json
// @Produce json
// @Param UpdateAccount body body.UpdateAccount true "Update Account data"
// @Success 200 {object} model.Account
// @Router /account [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.UpdateAccount

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := account.ID
	if id == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Account id is not valid.")
		return
	}

	if data.FCMToken == "" || data.InterestedField == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included.")
		return
	}

	if !data.ValidateInterestedFieldAttribute() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(
			httpError.Code(),
			"InterestedField is must be one of 'develop', 'design' and 'manage'.",
		)
		return
	}

	command := &command.UpdateCommand{
		AccountID:             id,
		Password:              data.Password,
		InterestedField:       data.InterestedField,
		FCMToken:              data.FCMToken,
		InterestedFieldDetail: data.InterestedFieldDetail,
	}

	updatedAccount, handlingError := controller.commandBus.Handle(command)

	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, updatedAccount)
}
