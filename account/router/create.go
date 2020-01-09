package router

import (
	"net/http"

	"github.com/kyhsa93/go-rest-example/account/dto"

	"github.com/gin-gonic/gin"
)

// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts [post]
func (router *Router) create(context *gin.Context) {
	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	duplicated := router.service.Create(&data)

	if duplicated != nil {
		httpError := router.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, "Account created")
}
