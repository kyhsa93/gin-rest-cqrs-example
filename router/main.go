package router

import "study/repositories"

import "github.com/gin-gonic/gin"

import "study/controller"

import "study/model"

func SetupRoutes(studyRepository *repositories.StudyRepository) *gin.Engine {
	route := gin.Default()

	route.POST("studies", func(context *gin.Context) {
		var study model.Study
		context.ShouldBindJSON(&study)
		controller.CreateStudy(&study, *studyRepository)
	})
	return route
}
