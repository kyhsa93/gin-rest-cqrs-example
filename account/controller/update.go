package controller

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/dto"
)

// @Description update account
// @Tags Accounts
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "accountId"
// @Param email formData string true "account email address"
// @Param provider formData string true "login service provider"
// @Param gender formData string true "user's gender male or female"
// @Param interest formData string true "interested part in develop, design, manage"
// @Param social_id formData string false "socialId when use social login"
// @Param password formData string false "need if don't use social login"
// @Param image formData file false "Profile image file"
// @Success 200
// @Router /accounts/{id} [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	controller.AuthenticateHTTPRequest(context)

	id := context.Param("id")
	email := context.PostForm("email")
	provider := context.PostForm("provider")
	socialID := context.PostForm("social_id")
	password := context.PostForm("password")
	gender := context.PostForm("gender")
	interest := context.PostForm("interest")

	if email == "" || provider == "" || gender == "" || interest == "" || (socialID == "" && password == "") {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Empty data is included.")
		return
	}

	data := dto.Account{
		Email:    email,
		Provider: provider,
		SocialID: socialID,
		Password: password,
		Gender:   gender,
		Interest: interest,
	}

	if !emailAndProviderValidation(data.Email, data.Provider) {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email and Provider is not matched.")
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email format is not valid.")
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Email host is not existed.")
		return
	}

	_, existedProvider := dto.Provider()[data.Provider]
	if existedProvider == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Provider is must one of 'email' or 'gmail'.")
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	if validate := dto.ValidateAccountAttributeByProvider(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if validate := dto.ValidateInterestAttribute(&data); validate == false {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "Interest is must be one of 'develop', 'design' and 'manage'.")
		return
	}

	image, _ := context.FormFile("image")

	command := &command.UpdateCommand{
		AccountID: id,
		Email:     email,
		Provider:  provider,
		SocialID:  socialID,
		Password:  password,
		Gender:    gender,
		Interest:  interest,
		Image:     image,
	}

	if err := controller.commandBus.Handle(command); err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, "Account updated")
}
