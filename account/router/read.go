package router

import (
	"github.com/gin-gonic/gin"
)

// ReadItem read account route handler
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account ID"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
func (router *Router) readItem(context *gin.Context) {
	context.JSON(200, router.service.ReadItem(context.Param("id")))
}

// ReadList read accounts route handler
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Accounts
// @Router /accounts [get]
func (router *Router) readList(context *gin.Context) {
	context.JSON(200, router.service.ReadList())
}
