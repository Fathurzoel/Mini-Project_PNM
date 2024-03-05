// app/controllers/quizController.go

package controllers

import (
	"WebDev/app/models"
	"WebDev/app/services"
	"WebDev/app/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QuizController bertanggung jawab untuk menangani permintaan terkait kuis.
type QuizController struct {
	QuizService  *services.QuizService
	ErrorHandler *utils.ErrorHandler
}

// NewQuizController menginisialisasi dan mengembalikan instance QuizController.
func NewQuizController(quizService *services.QuizService, errorHandler *utils.ErrorHandler) *QuizController {
	return &QuizController{
		QuizService:  quizService,
		ErrorHandler: errorHandler,
	}
}

// GetActiveQuizzesHandler menangani permintaan untuk mendapatkan daftar kuis aktif.
func (qc *QuizController) GetActiveQuizzesHandler(c *gin.Context) {
	activeQuizzes, err := qc.QuizService.GetActiveQuizzes()
	if err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar kuis aktif"})
		return
	}

	c.JSON(http.StatusOK, activeQuizzes)
}

// GetCompletedQuizzesHandler menangani permintaan untuk mendapatkan daftar kuis yang sudah selesai.
func (qc *QuizController) GetCompletedQuizzesHandler(c *gin.Context) {
	completedQuizzes, err := qc.QuizService.GetCompletedQuizzes()
	if err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar kuis selesai"})
		return
	}

	c.JSON(http.StatusOK, completedQuizzes)
}

// SearchQuizzesHandler menangani permintaan pencarian kuis berdasarkan judul.
func (qc *QuizController) SearchQuizzesHandler(c *gin.Context) {
	searchTerm := c.Query("q")

	if searchTerm == "" {
		qc.ErrorHandler.HandleError(nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pencarian tidak boleh kosong"})
		return
	}

	searchResults, err := qc.QuizService.SearchQuizzes(searchTerm)
	if err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan pencarian kuis"})
		return
	}

	c.JSON(http.StatusOK, searchResults)
}

// CreateQuizHandler menangani permintaan untuk membuat kuis baru.
func (qc *QuizController) CreateQuizHandler(c *gin.Context) {
	var newQuiz models.Quiz
	if err := c.ShouldBindJSON(&newQuiz); err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid"})
		return
	}

	if err := qc.QuizService.CreateQuiz(&newQuiz); err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat kuis"})
		return
	}

	c.JSON(http.StatusCreated, newQuiz)
}

// UpdateQuizHandler menangani permintaan untuk memperbarui kuis yang sudah ada.
func (qc *QuizController) UpdateQuizHandler(c *gin.Context) {
	quizID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID kuis tidak valid"})
		return
	}

	var updatedQuiz models.Quiz
	if err := c.ShouldBindJSON(&updatedQuiz); err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid"})
		return
	}

	updatedQuiz.ID = uint(quizID)

	if err := qc.QuizService.UpdateQuiz(&updatedQuiz); err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui kuis"})
		return
	}

	c.JSON(http.StatusOK, updatedQuiz)
}

// DeleteQuizHandler menangani permintaan untuk menghapus kuis.
func (qc *QuizController) DeleteQuizHandler(c *gin.Context) {
	quizID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID kuis tidak valid"})
		return
	}

	if err := qc.QuizService.DeleteQuiz(uint(quizID)); err != nil {
		qc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kuis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Kuis dengan ID %d telah dihapus", quizID)})
}
