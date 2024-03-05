// quizRoutes.go

package routes

import (
	"WebDev/app/controllers"
	"WebDev/app/services"
	"WebDev/app/utils"

	"github.com/gin-gonic/gin"
)

// InitializeQuizRoutes menginisialisasi rute-rute terkait kuis.
func InitializeQuizRoutes(router *gin.Engine, quizService *services.QuizService, errorHandler *utils.ErrorHandler) {
	quizController := controllers.NewQuizController(quizService, errorHandler)

	quizRoutes := router.Group("/quizzes")
	{
		quizRoutes.GET("/active", quizController.GetActiveQuizzesHandler)
		quizRoutes.GET("/completed", quizController.GetCompletedQuizzesHandler)
		quizRoutes.GET("/search", quizController.SearchQuizzesHandler)
		quizRoutes.POST("/", quizController.CreateQuizHandler)
		quizRoutes.PUT("/:id", quizController.UpdateQuizHandler)
		quizRoutes.DELETE("/:id", quizController.DeleteQuizHandler)
	}
}
