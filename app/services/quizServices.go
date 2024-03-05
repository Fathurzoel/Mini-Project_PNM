// backend/app/services/quizService.go

package services

import (
	"WebDev/app/config"
	"WebDev/app/models"
	"WebDev/app/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

// QuizService menyediakan layanan terkait quiz.
type QuizService struct {
	DB          *gorm.DB
	Error       *utils.ErrorHandler
	Config      *config.Config              // Menambahkan konfigurasi sebagai bagian dari service
	ErrorHandle utils.ErrorHandlerInterface // Menambahkan ErrorHandler sebagai bagian dari service
}

// NewQuizService menginisialisasi instance baru dari QuizService.
func NewQuizService(config *config.Config, db *gorm.DB, errorHandler *utils.ErrorHandler) *QuizService {
	return &QuizService{
		DB:          db,
		Error:       errorHandler,
		Config:      config,
		ErrorHandle: errorHandler, // Menggunakan ErrorHandler sebagai bagian dari service
	}
}

// GetActiveQuizzes mengambil quiz yang aktif dari database.
func (s *QuizService) GetActiveQuizzes() ([]models.Quiz, error) {
	// Mendapatkan waktu saat ini
	currentTime := time.Now()

	// Mengambil daftar quiz yang aktif dari database
	activeQuizzes := make([]models.Quiz, 0)
	err := s.DB.Where("start <= ? AND finish >= ?", currentTime, currentTime).Find(&activeQuizzes).Error
	if err != nil {
		return nil, err
	}

	return activeQuizzes, nil
}

// GetCompletedQuizzes mengambil quiz yang sudah selesai dari database.
func (s *QuizService) GetCompletedQuizzes() ([]models.Quiz, error) {
	// Mendapatkan waktu saat ini
	currentTime := time.Now()

	// Mengambil daftar quiz yang sudah selesai dari database
	completedQuizzes := make([]models.Quiz, 0)
	err := s.DB.Where("finish < ?", currentTime).Find(&completedQuizzes).Error
	if err != nil {
		return nil, err
	}

	return completedQuizzes, nil
}

// SearchQuizzes performs a search for quizzes based on a search term.
func (s *QuizService) SearchQuizzes(searchTerm string) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	// Menggunakan LOWER untuk pencarian yang tidak peka terhadap huruf besar/kecil
	err := s.DB.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(searchTerm)+"%").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

// GetQuizByID mengambil quiz berdasarkan ID-nya.
func (s *QuizService) GetQuizByID(quizID uint) (*models.Quiz, error) {
	// Mengambil quiz dari database berdasarkan ID
	quiz := &models.Quiz{}
	err := s.DB.First(quiz, quizID).Error
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

// CreateQuiz membuat quiz baru di database.
func (s *QuizService) CreateQuiz(quiz *models.Quiz) error {
	// Menyimpan quiz baru ke database
	err := s.DB.Create(quiz).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateQuiz memperbarui quiz yang sudah ada di database.
func (s *QuizService) UpdateQuiz(quiz *models.Quiz) error {
	// Menyimpan perubahan pada quiz ke database
	err := s.DB.Save(quiz).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteQuiz menghapus quiz dari database.
func (s *QuizService) DeleteQuiz(quizID uint) error {
	// Menghapus quiz berdasarkan ID-nya dari database
	err := s.DB.Where("id = ?", quizID).Delete(&models.Quiz{}).Error
	if err != nil {
		return err
	}

	return nil
}
