package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/model"
)

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account id"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
// @Security AccessToken
// @Security RefreshToken
func (router *Router) readAccountByID(context *gin.Context) {
	accessHeader := context.GetHeader("Authorization")

	if accessHeader == "" {
		httpError := router.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := context.Param("id")

	if id == "" {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	token := &model.Token{ID: id, Access: accessHeader}

	if auth := token.Validate(); auth == "" || auth != id {
		httpError := router.util.Error.HTTP.Forbidden()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	account := router.service.ReadAccountByID(id)

	if account == nil {
		httpError := router.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, account)
}

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Token
// @Router /accounts [get]
// @Param email query string true "account email"
// @Param provider query string true "account service provider"
// @Param password query string false "account password (email provider only)"
// @Param social_id query string false "account social_id"
func (router *Router) readAccount(context *gin.Context) {
	email := context.Query("email")
	socialID := context.Query("social_id")
	provider := context.Query("provider")
	password := context.Query("password")

	socialIDAndPasswordBothEmpty := false
	if socialID == "" && password == "" {
		socialIDAndPasswordBothEmpty = true
	}

	if email == "" || provider == "" || socialIDAndPasswordBothEmpty {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	data := &dto.Account{Email: email, Provider: provider, SocialID: socialID, Password: password}

	_, existedProvider := dto.Provider()[data.Provider]

	if existedProvider == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	dto.FilterAccountAttributeByProvider(data)

	validate := dto.ValidateAccountAttributeByProvider(data)
	if validate == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	account, err := router.service.ReadAccount(email, provider, socialID, password, false)

	if account == nil || err != nil {
		httpError := router.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	accessToken := account.CreateAccessToken()

	if accessToken == "" {
		httpError := router.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, &model.Token{ID: account.ID, Access: accessToken})
}
