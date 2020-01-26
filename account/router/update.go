package router

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/dto"
	"github.com/kyhsa93/gin-rest-example/account/model"
)

// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @param id path string true "account Id"
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts/{id} [put]
// @Security AccessToken
// @Security RefreshToken
func (router *Router) update(context *gin.Context) {
	accessHeader := context.GetHeader("Authorization")

	if accessHeader == "" {
		httpError := router.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := context.Param("id")

	if id == "" {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	token := &model.Token{ID: id, Access: accessHeader}

	if auth := token.Validate(); auth == "" || auth != id {
		httpError := router.util.Error.HTTP.Forbidden()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	_, existedProvider := dto.Provider()[data.Provider]

	if existedProvider == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	validate := dto.ValidateAccountAttributeByProvider(&data)
	if validate == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	router.service.Update(id, data.Email, data.Provider, data.SocialID, data.Password)

	context.JSON(http.StatusOK, "Account updated")
}
