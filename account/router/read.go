package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/model"
)

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account id"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
func (router *Router) readAccount(context *gin.Context) {
	accessHeader := context.GetHeader("Authorization")
	refreshHeader := context.GetHeader("Refresh")

	if accessHeader == "" || refreshHeader == "" {
		httpError := router.util.HTTPError.Unauthorized()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	id := context.Param("id")

	if id == "" {
		httpError := router.util.HTTPError.BadRequest()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	token := &model.Token{ID: id, Access: accessHeader, Refresh: refreshHeader}

	if auth := token.Validate(); auth == "" || auth != id {
		httpError := router.util.HTTPError.Forbidden()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	account := router.service.ReadAccount(id)

	if account == nil {
		httpError := router.util.HTTPError.NotFound()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	context.JSON(200, account)
}

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Token
// @Router /accounts [get]
// @Param email query string true "account email"
// @Param social_id query string true "account social_id"
func (router *Router) readAccountByEmailAndSocialID(context *gin.Context) {
	email := context.Query("email")
	socialID := context.Query("social_id")

	if email == "" || socialID == "" {
		httpError := router.util.HTTPError.BadRequest()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	account := router.service.ReadAccountByEmailAndSocialID(email, socialID)

	if account == nil {
		httpError := router.util.HTTPError.NotFound()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	accessToken := account.CreateAccessToken()

	if accessToken == "" {
		httpError := router.util.HTTPError.Error()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	refreshToken := account.CreateRefreshToken(accessToken)

	if refreshToken == "" {
		httpError := router.util.HTTPError.Error()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	context.JSON(200, &model.Token{ID: account.ID, Access: accessToken, Refresh: refreshToken})
}
