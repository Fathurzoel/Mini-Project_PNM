// main.go

package main

import (
	"WebDev/app/config"
	"WebDev/app/controllers"
	"WebDev/app/middleware"
	"WebDev/app/routes"
	"WebDev/app/services"
	"WebDev/app/utils"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Memuat konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Gagal memuat konfigurasi:", err)
	}

	// Menginisialisasi koneksi database
	db, err := config.InitializeDBConnection()
	if err != nil {
		log.Fatal("Gagal menginisialisasi database:", err)
	}

	// Menginisialisasi layanan-layanan
	errorHandler := utils.NewErrorHandler("error", log.New(gin.DefaultWriter, "\r\n", 0), nil)
	authService := services.NewAuthService(db, errorHandler, cfg)

	// Menginisialisasi pengontrol-pengontrol
	authController := controllers.NewAuthController(authService, errorHandler)

	// Menginisialisasi router
	router := gin.Default()

	// Menambahkan middleware CORS
	router.Use(cors.Default())

	// Menambahkan middleware untuk verifikasi token JWT
	router.Use(middleware.AuthMiddleware(db, errorHandler, cfg))

	// Menginisialisasi rute-rute
	routes.SetAuthRoutes(router, authController)

	// Menjalankan server
	port := ":8080"
	fmt.Println("Server berjalan di port", port)
	err = router.Run(port)
	if err != nil {
		log.Fatal("Gagal memulai server:", err)
	}
}
