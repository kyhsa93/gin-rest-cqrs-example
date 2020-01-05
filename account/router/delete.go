package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200
// @Router /accounts/{id} [delete]
func (router *Router) delete(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		httpError := router.util.HTTPError.BadRequest()
		context.JSON(httpError.Code, httpError.Message)
		return
	}

	router.service.Delete(id)

	context.JSON(http.StatusOK, "Account deleted")
}
