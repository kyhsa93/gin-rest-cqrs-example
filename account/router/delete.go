package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/controller"
)

// Delete delete account route handler
// @Description delete account by id
// @Tags Accounts
// @Param id path string true "account ID"
// @Success 200
// @Router /accounts/{id} [delete]
func Delete(context *gin.Context) {
	controller.Delete(context.Param("id"))
}
