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
// @Param CreateAccount body body.CreateAccount true "Create Account data"
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

	query := &query.ReadAccountByEmailQuery{Email: data.Email}
	duplicated, _ := controller.queryBus.Handle(query)
	if duplicated.ID != "" {
		httpError := controller.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), "Email is duplicated.")
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	if !dto.ValidateAccountAttributeByProvider(&data) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if !dto.ValidateInterestedFieldAttribute(&data) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(
			httpError.Code(),
			"InterestedField is must be one of 'develop', 'design' and 'manage'.",
		)
		return
	}

	command := &command.CreateCommand{
		Email:                 data.Email,
		Provider:              data.Provider,
		SocialID:              data.SocialID,
		Password:              data.Password,
		Gender:                data.Gender,
		InterestedField:       data.InterestedField,
		InterestedFieldDetail: data.InterestedFieldDetail,
	}

	createdAccount, handlingError := controller.commandBus.Handle(command)
	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, createdAccount)
}
