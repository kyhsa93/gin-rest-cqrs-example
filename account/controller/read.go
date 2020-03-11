package controller

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/dto"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"
)

// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} model.Account
// @Router /accounts [get]
// @Param email query string false "account email"
// @Param provider query string false "account service provider"
// @Param password query string false "account password (email provider only)"
// @Param social_id query string false "account social_id"
// @Security AccessToken
func (controller *Controller) readAccount(context *gin.Context) {
	email := context.Query("email")
	socialID := context.Query("social_id")
	provider := context.Query("provider")
	password := context.Query("password")

	if email == "" && socialID == "" && provider == "" && password == "" {
		accessToken := context.GetHeader("Authorization")
		account, err := controller.GetAccountByAccessToken(accessToken)
		if err != nil {
			httpError := controller.util.Error.HTTP.Unauthorized()
			context.JSON(httpError.Code(), httpError.Message())
			return
		}
		context.JSON(http.StatusOK, account)
		return
	}

	socialIDAndPasswordBothEmpty := false
	if socialID == "" && password == "" {
		socialIDAndPasswordBothEmpty = true
	}

	if email == "" || provider == "" || socialIDAndPasswordBothEmpty {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(email)
	if emaiFormatlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(email)
	if emaiHostlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email host is not existed.")
		return
	}

	data := &dto.ReadAccount{
		Email:    email,
		Provider: provider,
		SocialID: socialID,
		Password: password,
	}

	if !data.ValidateProvider() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Provider is must one of 'email' or 'gmail'.")
		return
	}

	data.SetAccountAttributeByProvider()

	if !data.ValidateAccountAttributeByProvider() {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadAccountQuery{
		Email:    email,
		Provider: provider,
		SocialID: socialID,
		Password: password,
		Unscoped: false,
	}
	account, _ := controller.queryBus.Handle(query)
	if account == nil {
		context.JSON(http.StatusOK, account)
		return
	}

	context.JSON(http.StatusOK, account)
}
