package controller

import "github.com/gin-gonic/gin"

func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	accountID := context.PostForm("account_id")
	auth := controller.AuthenticateHTTPReqeust(accessToken, accountID)
	if !auth {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
}
