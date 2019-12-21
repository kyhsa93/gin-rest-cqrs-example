package router

import (
	"github.com/kyhsa93/go-rest-example/account/dto"

	"github.com/gin-gonic/gin"
)

// Create create account route handler
// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts [post]
func (router *Router) create(context *gin.Context) {
	var data dto.Account
	context.ShouldBindJSON(&data)
	router.service.Create(&data)
}
