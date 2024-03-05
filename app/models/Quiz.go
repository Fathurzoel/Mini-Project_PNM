// backend/app/models/quiz.go

package models

import (
	"time"
)

// Quiz model mewakili entitas kuis.
type Quiz struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Start       time.Time  `json:"start"`
	Finish      time.Time  `json:"finish"`
	Questions   []Question `gorm:"foreignKey:QuizID" json:"questions"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
}

// Question model mewakili entitas pertanyaan.
type Question struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	QuizID        uint      `json:"-"`
	QuestionText  string    `gorm:"type:text" json:"question_text"`
	OptionA       string    `gorm:"size:255" json:"option_a"`
	OptionB       string    `gorm:"size:255" json:"option_b"`
	OptionC       string    `gorm:"size:255" json:"option_c"`
	OptionD       string    `gorm:"size:255" json:"option_d"`
	CorrectAnswer string    `gorm:"size:1" json:"correct_answer"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
