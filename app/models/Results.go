// models/results.go

package models

import "gorm.io/gorm"

// Participant model mewakili tabel participant dalam basis data.
type Participant struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	QuizID uint   `json:"quiz_id"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

// ParticipantAnswer model mewakili tabel participant_answer dalam basis data.
type ParticipantAnswer struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	ParticipantID uint   `json:"participant_id"`
	QuestionID    uint   `json:"question_id"`
	Answer        string `json:"answer"`
}

// BeforeCreate hook untuk model Participant.
func (p *Participant) BeforeCreate(tx *gorm.DB) error {
	return nil // Implementasikan jika diperlukan
}

// BeforeCreate hook untuk model ParticipantAnswer.
func (pa *ParticipantAnswer) BeforeCreate(tx *gorm.DB) error {
	return nil // Implementasikan jika diperlukan
}
