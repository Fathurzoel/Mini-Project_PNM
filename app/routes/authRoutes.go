// app/routes/authRoutes.go

package routes

import (
	"WebDev/app/controllers"

	"github.com/gin-gonic/gin"
)

// SetAuthRoutes sets up authentication-related routes.
func SetAuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.LoginHandler)
		// Add other authentication-related routes as needed
	}
}
