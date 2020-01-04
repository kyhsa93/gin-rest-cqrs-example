package router

import (
	"github.com/kyhsa93/go-rest-example/account/dto"

	"github.com/gin-gonic/gin"
)

// Update update route handler
// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @param id path string true "account Id"
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts/{id} [put]
func (router *Router) update(context *gin.Context) {
	var data dto.Account
	context.ShouldBindJSON(&data)
	router.service.Update(context.Param("id"), &data)
}
