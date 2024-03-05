// backend/app/models/user.go

package models

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// User mewakili entitas Pengguna dalam basis data.
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Nama     string `json:"nama"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	JWTToken string `json:"jwt_token"`
}

// LoginRequest mewakili struktur permintaan untuk login pengguna.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Validate melakukan validasi pada LoginRequest.
func (lr *LoginRequest) Validate() error {
	if lr.Email == "" || lr.Password == "" {
		return errors.New("email dan password diperlukan")
	}

	if !strings.Contains(lr.Email, "@") {
		return fmt.Errorf("email harus berupa alamat email yang valid")
	}

	return nil
}

// AutoMigrateModels melakukan migrasi otomatis untuk model yang sudah ada.
func AutoMigrateModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

// InitializeModels menginisialisasi model dan melakukan migrasi.
func InitializeModels(db *gorm.DB) {
	AutoMigrateModels(db)
}
