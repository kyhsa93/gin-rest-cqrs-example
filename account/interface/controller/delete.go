package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/domain/model"
)

// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200
// @Router /accounts/{id} [delete]
// @Security AccessToken
// @Security RefreshToken
func (controller *Controller) delete(context *gin.Context) {
	accessHeader := context.GetHeader("Authorization")

	if accessHeader == "" {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := context.Param("id")

	if id == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	token := &model.Token{ID: id, Access: accessHeader}

	if auth := token.Validate(); auth == "" || auth != id {
		httpError := controller.util.Error.HTTP.Forbidden()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	controller.application.Delete(id)

	context.JSON(http.StatusOK, "Account deleted")
}
