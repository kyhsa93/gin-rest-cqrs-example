package router

import (
	"github.com/gin-gonic/gin"
)

// ReadAccount read account route handler
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account id"
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
// @Param social_id query string true "account social_id"
func (router *Router) readAccounts(context *gin.Context) {
	email := context.Query("email")
	socialID := context.Query("social_id")
	context.JSON(200, router.service.ReadAccounts(email, socialID))
}
