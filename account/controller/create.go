package controller

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/dto"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"

	"github.com/gin-gonic/gin"
)

// @Description create account
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account body dto.Account true "Create Account data"
// @Success 201 {object} model.Account
// @Router /accounts [post]
func (controller *Controller) create(context *gin.Context) {
	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if data.Email == "" || data.Provider == "" || data.Gender == "" ||
		data.Interest == "" || (data.SocialID == "" && data.Password == "") {
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

	query := &query.ReadAccountQuery{
		Email: data.Email, Password: "", Provider: "", SocialID: "", Unscoped: true,
	}
	duplicated, _ := controller.queryBus.Handle(query)
	if duplicated != nil {
		httpError := controller.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), "Email is duplicated.")
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	if validate := dto.ValidateAccountAttributeByProvider(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if validate := dto.ValidateInterestAttribute(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Interest is must be one of 'develop', 'design' and 'manage'.")
		return
	}

	command := &command.CreateCommand{
		Email:    data.Email,
		Provider: data.Provider,
		SocialID: data.SocialID,
		Password: data.Password,
		Gender:   data.Gender,
		Interest: data.Interest,
		FileID:   "",
	}

	createdAccount, hadlingError := controller.commandBus.Handle(command)
	if hadlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, createdAccount)
}
