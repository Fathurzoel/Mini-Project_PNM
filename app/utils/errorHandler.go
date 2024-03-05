// app/services/errorHandler.go

package utils

import (
	"log"
)

// ErrorHandlerInterface menyediakan metode HandleError.
type ErrorHandlerInterface interface {
	HandleError(err error) error
}

// ErrorResponse mewakili struktur untuk respons kesalahan.
type ErrorResponse struct {
	Message string `json:"message"`
	// Anda bisa menambahkan properti lain sesuai kebutuhan, misalnya Code untuk kode kesalahan.
}

// Error mengimplementasikan antarmuka error.
func (e *ErrorResponse) Error() string {
	panic("belum diimplementasikan")
}

// ErrorHandler menangani kesalahan dan mengembalikan respons yang sesuai.
type ErrorHandler struct {
	LogLevel    string // LogLevel bisa berisi nilai seperti "info", "warn", atau "error"
	Logger      *log.Logger
	AuthService ErrorHandlerInterface // Menggunakan antarmuka ErrorHandlerInterface
}

// NewErrorHandler menginisialisasi instance baru dari ErrorHandler.
func NewErrorHandler(logLevel string, logger *log.Logger, authService ErrorHandlerInterface) *ErrorHandler {
	return &ErrorHandler{
		LogLevel:    logLevel,
		Logger:      logger,
		AuthService: authService,
	}
}

// HandleError menangani kesalahan dan mengembalikan respons yang sesuai.
func (e *ErrorHandler) HandleError(err error) error {
	// Catat kesalahan berdasarkan tingkat log
	switch e.LogLevel {
	case "info":
		e.Logger.Println("INFO:", err.Error())
	case "warn":
		e.Logger.Println("WARNING:", err.Error())
	case "error":
		e.Logger.Println("ERROR:", err.Error())
	default:
		// Default untuk mencatat sebagai kesalahan
		e.Logger.Println("ERROR:", err.Error())
	}

	// Menggunakan metode dari AuthService
	return e.AuthService.HandleError(err)
}
