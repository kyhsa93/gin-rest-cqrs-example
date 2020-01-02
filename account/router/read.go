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
// @Success 200 {object} model.Account
// @Router /accounts [get]
// @Param email query string true "account email"
// @Param password query string true "account password"
func (router *Router) readAccounts(context *gin.Context) {
	email := context.Query("email")
	password := context.Query("password")
	context.JSON(200, router.service.ReadAccounts(email, password))
}
