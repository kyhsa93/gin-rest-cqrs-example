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
// @param id path string true "account Id"
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts/{id} [put]
func (router *Router) update(context *gin.Context) {
	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := router.util.HTTPError.BadRequest()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	router.service.Update(context.Param("id"), &data)

	context.JSON(http.StatusOK, "Account updated")
}
