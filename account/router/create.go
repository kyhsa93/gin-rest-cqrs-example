package router

import (
	"net/http"

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
	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse body"})
	}

	if validationError := data.Validate(&data); validationError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": validationError.Error()})
	}
	router.service.Create(&data)
}
