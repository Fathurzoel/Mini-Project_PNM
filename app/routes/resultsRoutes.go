// backend/app/routes/resultsRoutes.go

package routes

import (
	"WebDev/app/controllers"
	"WebDev/app/services"
	"WebDev/app/utils"

	"github.com/gin-gonic/gin"
)

// SetResultsRoutes mengatur rute-rute terkait hasil kuis.
func SetResultsRoutes(router *gin.RouterGroup, resultsService *services.ResultsService, errorHandler *utils.ErrorHandler) {
	resultsController := controllers.NewResultsController(resultsService, errorHandler)

	router.GET("/quizzes/:id/results", resultsController.GetResultsByQuizIDHandler)
	router.GET("/participants/:id", resultsController.GetParticipantDetailsHandler)
	router.POST("/results", resultsController.AddParticipantResultHandler)
	router.PUT("/participants/:id/results", resultsController.UpdateParticipantResultHandler)
	router.DELETE("/participants/:id/results", resultsController.DeleteParticipantResultHandler)
}
