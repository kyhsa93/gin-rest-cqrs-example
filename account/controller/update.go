package controller

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/dto"
)

// @Description update account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param id path string true "account id"
// @Param account body dto.Account true "Update Account data"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	controller.AuthenticateHTTPRequest(context)
	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := context.Param("id")
	if id == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Account id is not valid.")
		return
	}

	if data.Email == "" || data.Provider == "" || data.Gender == "" ||
		data.InterestedField == "" || (data.SocialID == "" && data.Password == "") {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included.")
		return
	}

	if !emailAndProviderValidation(data.Email, data.Provider) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email and Provider is not matched.")
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email format is not valid.")
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email host is not existed.")
		return
	}

	_, existedProvider := dto.Provider()[data.Provider]
	if existedProvider == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Provider is must one of 'email' or 'gmail'.")
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	if validate := dto.ValidateAccountAttributeByProvider(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if validate := dto.ValidateInterestedFieldAttribute(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "InterestedField is must be one of 'develop', 'design' and 'manage'.")
		return
	}

	if !controller.ValidateFileID(id, data.FileID) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "FileId is not matched.")
		return
	}

	command := &command.UpdateCommand{
		AccountID:       id,
		Email:           data.Email,
		Provider:        data.Provider,
		SocialID:        data.SocialID,
		Password:        data.Password,
		Gender:          data.Gender,
		InterestedField: data.InterestedField,
		FileID:          data.FileID,
	}

	updatedAccount, handlingError := controller.commandBus.Handle(command)

	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, updatedAccount)
}
