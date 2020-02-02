package router

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/kyhsa93/gin-rest-example/account/dto"

	"github.com/gin-gonic/gin"
)

// @Description create account
// @Tags Accounts
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "account email address"
// @Param provider formData string true "login service provider"
// @Param gender formData string true "user's gender male of female"
// @Param image formData file false "Profile image file"
// @Param social_id formData string false "socialId when use social login"
// @Param password formData string false "need if don't use social login"
// @Success 201
// @Router /accounts [post]
func (router *Router) create(context *gin.Context) {
	email := context.PostForm("email")
	provider := context.PostForm("provider")
	socialID := context.PostForm("social_id")
	password := context.PostForm("password")
	gender := context.PostForm("gender")

	if email == "" || provider == "" || gender == "" || (socialID == "" && password == "") {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	data := dto.Account{
		Email:    email,
		Provider: provider,
		SocialID: socialID,
		Password: password,
		Gender:   gender,
	}

	if !emailAndProviderValidation(data.Email, data.Provider) {
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

	duplicated, _ := router.service.ReadAccount(data.Email, "", "", "", true)

	if duplicated != nil {
		httpError := router.util.Error.HTTP.Conflict()
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

	image, _ := context.FormFile("image")

	router.service.Create(data.Email, data.Provider, data.SocialID, data.Password, image, data.Gender)

	context.JSON(http.StatusCreated, "Account created")
}
