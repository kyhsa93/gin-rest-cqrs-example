package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/controller"
	"github.com/kyhsa93/go-rest-example/account/repository"
)

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param id path string true "account ID"
// @Success 200 {object} model.Account
// @Router /accounts/{id} [get]
func ReadItem(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadItem(context.Param("id"), repository))
}

// @Tags Accounts
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Accounts
// @Router /accounts [get]
func ReadList(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadList(repository))
}
