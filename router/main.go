package router

import (
	"study/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(studyRepository *repositories.StudyRepository) *gin.Engine {
	route := gin.Default()

	route.POST("studies", func(context *gin.Context) {
		Create(context, *studyRepository)
	})

	route.GET("studies/:id", func(context *gin.Context) {
		Read(context, *studyRepository)
	})

	route.GET("studies", func(context *gin.Context) {
		ReadList(context, *studyRepository)
	})

	route.PUT("studies/:id", func(context *gin.Context) {
		Update(context, *studyRepository)
	})

	route.DELETE("studies/:id", func(context *gin.Context) {
		Delete(context, *studyRepository)
	})

	return route
}
