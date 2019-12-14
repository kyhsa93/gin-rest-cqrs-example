package router

import (
	"github.com/kyhsa93/go-rest-example/account/controller"
	"github.com/kyhsa93/go-rest-example/account/dto"
	"github.com/kyhsa93/go-rest-example/account/repository"

	"github.com/gin-gonic/gin"
)

// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @param id path string true "account ID"
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts/{id} [put]
func Update(context *gin.Context, repository repository.Repository) {
	var data dto.Account
	context.ShouldBindJSON(&data)
	controller.Update(context.Param("id"), &data, repository)
}
