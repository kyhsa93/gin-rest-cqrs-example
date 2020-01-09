package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/model"
)

// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200
// @Router /accounts/{id} [delete]
func (router *Router) delete(context *gin.Context) {
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

	router.service.Delete(id)

	context.JSON(http.StatusOK, "Account deleted")
}
