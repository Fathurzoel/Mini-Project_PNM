// backend/app/controllers/resultsController.go

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

// ResultsController bertanggung jawab untuk menangani permintaan terkait hasil kuis.
type ResultsController struct {
	ResultsService *services.ResultsService
	ErrorHandler   *utils.ErrorHandler
}

// NewResultsController menginisialisasi dan mengembalikan instance ResultsController.
func NewResultsController(resultsService *services.ResultsService, errorHandler *utils.ErrorHandler) *ResultsController {
	return &ResultsController{
		ResultsService: resultsService,
		ErrorHandler:   errorHandler,
	}
}

// GetResultsByQuizIDHandler menangani permintaan untuk mendapatkan hasil kuis berdasarkan ID kuis.
func (rc *ResultsController) GetResultsByQuizIDHandler(c *gin.Context) {
	quizID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID kuis tidak valid"})
		return
	}

	results, err := rc.ResultsService.GetResultsByQuizID(uint(quizID))
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil hasil kuis"})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetParticipantDetailsHandler menangani permintaan untuk mendapatkan detail peserta termasuk jawaban mereka.
func (rc *ResultsController) GetParticipantDetailsHandler(c *gin.Context) {
	participantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID peserta tidak valid"})
		return
	}

	participant, err := rc.ResultsService.GetParticipantDetails(uint(participantID))
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil detail peserta"})
		return
	}

	c.JSON(http.StatusOK, participant)
}

// AddParticipantResultHandler menangani permintaan untuk menambahkan hasil kuis peserta baru.
func (rc *ResultsController) AddParticipantResultHandler(c *gin.Context) {
	var newResult models.Participant
	if err := c.ShouldBindJSON(&newResult); err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid"})
		return
	}

	if err := rc.ResultsService.AddParticipantResult(&newResult); err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan hasil kuis peserta"})
		return
	}

	c.JSON(http.StatusCreated, newResult)
}

// UpdateParticipantResultHandler menangani permintaan untuk memperbarui hasil kuis peserta.
func (rc *ResultsController) UpdateParticipantResultHandler(c *gin.Context) {
	participantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID peserta tidak valid"})
		return
	}

	var updatedResult models.Participant
	if err := c.ShouldBindJSON(&updatedResult); err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid"})
		return
	}

	updatedResult.ID = uint(participantID)

	if err := rc.ResultsService.UpdateParticipantResult(&updatedResult); err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui hasil kuis peserta"})
		return
	}

	c.JSON(http.StatusOK, updatedResult)
}

// DeleteParticipantResultHandler menangani permintaan untuk menghapus hasil kuis peserta.
func (rc *ResultsController) DeleteParticipantResultHandler(c *gin.Context) {
	participantID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID peserta tidak valid"})
		return
	}

	if err := rc.ResultsService.DeleteParticipantResult(uint(participantID)); err != nil {
		rc.ErrorHandler.HandleError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus hasil kuis peserta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hasil kuis peserta dengan ID %d telah dihapus", participantID)})
}
