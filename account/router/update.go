package router

import (
	"net/http"

	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/model"

	"github.com/gin-gonic/gin"
)

// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @param id path string true "account Id"
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts/{id} [put]
// @Security AccessToken
// @Security RefreshToken
func (router *Router) update(context *gin.Context) {
	accessHeader := context.GetHeader("Authorization")
	refreshHeader := context.GetHeader("Refresh")

	if accessHeader == "" || refreshHeader == "" {
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

	token := &model.Token{ID: id, Access: accessHeader, Refresh: refreshHeader}

	if auth := token.Validate(); auth == "" || auth != id {
		httpError := router.util.Error.HTTP.Forbidden()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	router.service.Update(id, &data)

	context.JSON(http.StatusOK, "Account updated")
}
