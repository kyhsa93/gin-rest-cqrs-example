package router

import (
	"github.com/gin-gonic/gin"
)

// ReadAccount read account route handler
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account ID"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
func (router *Router) readAccount(context *gin.Context) {
	context.JSON(200, router.service.ReadAccount(context.Param("id")))
}

// ReadAccounts read accounts route handler
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Accounts
// @Router /accounts [get]
func (router *Router) readAccounts(context *gin.Context) {
	context.JSON(200, router.service.ReadAccounts())
}
