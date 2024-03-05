//app/services/resultsService.go

package services

import (
	"WebDev/app/models"
	"WebDev/app/utils"

	"gorm.io/gorm"
)

// ResultsService menyediakan layanan terkait hasil kuis.
type ResultsService struct {
	DB    *gorm.DB
	Error *utils.ErrorHandler
}

// NewResultsService menginisialisasi instance baru dari ResultsService.
func NewResultsService(db *gorm.DB, errorHandler *utils.ErrorHandler) *ResultsService {
	return &ResultsService{
		DB:    db,
		Error: errorHandler,
	}
}

// GetResultsByQuizID mengambil hasil kuis berdasarkan ID kuis.
func (s *ResultsService) GetResultsByQuizID(quizID uint) ([]models.Participant, error) {
	var results []models.Participant

	err := s.DB.Where("quiz_id = ?", quizID).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetParticipantDetails mengambil detail peserta termasuk jawaban mereka.
func (s *ResultsService) GetParticipantDetails(participantID uint) (*models.Participant, error) {
	var participant models.Participant

	err := s.DB.Preload("ParticipantAnswers").First(&participant, participantID).Error
	if err != nil {
		return nil, err
	}

	return &participant, nil
}

// AddParticipantResult menambahkan hasil kuis peserta baru ke database.
func (s *ResultsService) AddParticipantResult(result *models.Participant) error {
	err := s.DB.Create(result).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateParticipantResult memperbarui hasil kuis peserta yang sudah ada di database.
func (s *ResultsService) UpdateParticipantResult(result *models.Participant) error {
	err := s.DB.Save(result).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteParticipantResult menghapus hasil kuis peserta dari database.
func (s *ResultsService) DeleteParticipantResult(participantID uint) error {
	err := s.DB.Where("id = ?", participantID).Delete(&models.Participant{}).Error
	if err != nil {
		return err
	}

	return nil
}
