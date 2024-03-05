// validators.go

package utils

import (
	"regexp"
	"unicode"
)

// IsValidEmail memeriksa apakah alamat email memiliki format yang benar.
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsStrongPassword memeriksa apakah kata sandi cukup kuat.
func IsStrongPassword(password string) bool {
	return len(password) >= 8 && containsUpperCase(password) && containsLowerCase(password) && containsDigit(password)
}

// IsValidUsername memeriksa apakah username sesuai dengan aturan tertentu.
func IsValidUsername(username string) bool {
	return len(username) >= 4
}

// containsUpperCase memeriksa apakah sebuah string mengandung huruf kapital.
func containsUpperCase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

// containsLowerCase memeriksa apakah sebuah string mengandung huruf kecil.
func containsLowerCase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

// containsDigit memeriksa apakah sebuah string mengandung angka.
func containsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
