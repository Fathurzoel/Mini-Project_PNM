// backend/app/middleware/authMiddleware.go

package middleware

import (
	"net/http"

	"WebDev/app/config"
	"WebDev/app/services"
	"WebDev/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthMiddleware adalah middleware untuk verifikasi token JWT.
func AuthMiddleware(db *gorm.DB, errorHandler *utils.ErrorHandler, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header
		tokenString := c.GetHeader("Authorization")

		// Jika tidak ada token, lanjutkan dengan rute tanpa verifikasi token
		if tokenString == "" {
			c.Next()
			return
		}

		// Verifikasi token jika token diberikan
		authService := services.NewAuthService(db, errorHandler, config)
		token, err := authService.VerifyToken(tokenString)

		// Jika terdapat kesalahan atau token kadaluwarsa, berikan respons Unauthorized
		if err != nil || authService.IsTokenExpired(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau telah kadaluwarsa"})
			c.Abort()
			return
		}

		// Setel ID pengguna dari klaim ke konteks, jika perlu
		email := int(token.Claims.(jwt.MapClaims)["user_id"].(float64))
		c.Set("email", email)

		c.Next()
	}
}
