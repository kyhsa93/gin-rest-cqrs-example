package router

import (
	"github.com/gin-gonic/gin"
)

// Delete delete account route handler
// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account Id"
// @Success 200
// @Router /accounts/{id} [delete]
func (router *Router) delete(context *gin.Context) {
	router.service.Delete(context.Param("id"))
}
