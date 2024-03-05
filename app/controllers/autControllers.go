// app/controllers/authController.go

package controllers

import (
	"WebDev/app/services"
	"WebDev/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController menangani permintaan HTTP terkait otentikasi.
type AuthController struct {
	AuthService  *services.AuthService
	ErrorHandler *utils.ErrorHandler
}

// NewAuthController menginisialisasi instance baru dari AuthController.
func NewAuthController(authService *services.AuthService, errorHandler *utils.ErrorHandler) *AuthController {
	return &AuthController{
		AuthService:  authService,
		ErrorHandler: errorHandler,
	}
}

// LoginRequest merepresentasikan struktur untuk permintaan login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse merepresentasikan struktur untuk respons login.
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginHandler menangani permintaan login.
func (c *AuthController) LoginHandler(ctx *gin.Context) {
	// Mengambil kredensial pengguna dari permintaan.
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	// Memeriksa apakah pengguna sudah memiliki token
	user, err := c.AuthService.FindUserByEmail(email)
	if err != nil {
		// Mengatasi kesalahan (mis., log)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan server internal"})
		return
	}

	if user != nil && user.JWTToken != "" {
		// Pengguna sudah memiliki token, tangani sesuai (mis., kembalikan kesalahan)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Pengguna sudah masuk"})
		return
	}

	// Memanggil AuthService untuk mengotentikasi pengguna.
	token, err := c.AuthService.Login(email, password)
	if err != nil {
		// Mengembalikan respons kesalahan ke klien.
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau kata sandi tidak valid"})
		return
	}

	// Mengembalikan respons yang sesuai berdasarkan hasil otentikasi.
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
