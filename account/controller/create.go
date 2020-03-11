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
// @Tags Account
// @Accept json
// @Produce json
// @Param CreateAccount body body.CreateAccount true "Create Account data"
// @Success 201 {object} model.Account
// @Router /accounts [post]
func (controller *Controller) create(context *gin.Context) {
	var data dto.CreateAccount

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if data.Email == "" || data.Provider == "" || data.Gender == "" ||
		data.InterestedField == "" || (data.SocialID == "" && data.Password == "") {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included")
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email format is not valid")
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email host is not existed")
		return
	}

	if !data.ValidateAccountGender() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Gender is nust one of 'male' or female")
		return
	}

	if !data.ValidateProvider() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Provider is must one of 'email' or 'gmail'")
		return
	}

	query := &query.ReadAccountByEmailQuery{Email: data.Email}
	duplicated, _ := controller.queryBus.Handle(query)
	if duplicated.ID != "" {
		httpError := controller.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), "Email is duplicated")
		return
	}

	data.SetAccountAttributeByProvider()

	if !data.ValidateAccountAttributeByProvider() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if !data.ValidateInterestedFieldAttribute() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(
			httpError.Code(),
			"InterestedField is must be one of 'develop', 'design' and 'manage'",
		)
		return
	}

	command := &command.CreateCommand{
		Email:                 data.Email,
		Provider:              data.Provider,
		SocialID:              data.SocialID,
		Password:              data.Password,
		Gender:                data.Gender,
		FCMToken:              data.FCMToken,
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
