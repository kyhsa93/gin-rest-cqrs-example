package controller

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/kyhsa93/gin-rest-example/account/application/command"
	"github.com/kyhsa93/gin-rest-example/account/application/query"
	"github.com/kyhsa93/gin-rest-example/account/interface/dto"

	"github.com/gin-gonic/gin"
)

// @Description create account
// @Tags Accounts
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "account email address"
// @Param provider formData string true "login service provider"
// @Param gender formData string true "user's gender male or female"
// @Param intereste formData string true "interested part in develop, design, manage"
// @Param social_id formData string false "socialId when use social login"
// @Param password formData string false "need if don't use social login"
// @Param image formData file false "Profile image file"
// @Success 201
// @Router /accounts [post]
func (controller *Controller) create(context *gin.Context) {
	email := context.PostForm("email")
	provider := context.PostForm("provider")
	socialID := context.PostForm("social_id")
	password := context.PostForm("password")
	gender := context.PostForm("gender")
	intereste := context.PostForm("intereste")

	if email == "" || provider == "" || gender == "" || intereste == "" || (socialID == "" && password == "") {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	data := dto.Account{
		Email:     email,
		Provider:  provider,
		SocialID:  socialID,
		Password:  password,
		Gender:    gender,
		Intereste: intereste,
	}

	if !emailAndProviderValidation(data.Email, data.Provider) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	_, existedProvider := dto.Provider()[data.Provider]
	if existedProvider == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadAccountQuery{
		Email: data.Email, Password: "", Provider: "", SocialID: "", Unscoped: true,
	}
	duplicated, _ := controller.queryBus.Handle(query)
	if duplicated != nil {
		httpError := controller.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	if validate := dto.ValidateAccountAttributeByProvider(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if validate := dto.ValidateInteresteAttribute(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	image, _ := context.FormFile("image")

	command := &command.CreateCommand{
		Email:     email,
		Provider:  provider,
		SocialID:  socialID,
		Password:  password,
		Gender:    gender,
		Intereste: intereste,
		Image:     image,
	}

	if err := controller.commandBus.Handle(command); err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, "Account created")
}
