package services

import (
	"WebDev/app/config"
	"WebDev/app/models"
	"WebDev/app/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthServiceInterface mendefinisikan antarmuka untuk AuthService.
type AuthServiceInterface interface {
	GenerateJWTToken(email string) (string, error)
	Login(email, password string) (string, error)
	FindUserByEmail(email string) (*models.User, error)
	FindAdminCredentials() (*models.User, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
	IsTokenExpired(token *jwt.Token) bool
	HandleError(err error) error
}

// AuthService menyediakan layanan otentikasi pengguna.
type AuthService struct {
	DB     *gorm.DB
	Config *config.Config
	Error  *utils.ErrorHandler
	User   *models.User
}

// NewAuthService menginisialisasi instance baru dari AuthService.
func NewAuthService(db *gorm.DB, errorHandler *utils.ErrorHandler, config *config.Config) *AuthService {
	return &AuthService{
		DB:     db,
		Error:  errorHandler,
		Config: config,
	}
}

// GenerateJWTToken menghasilkan token JWT berdasarkan email yang diberikan.
func (a *AuthService) GenerateJWTToken(email string) (string, error) {
	// Membuat objek token baru, menentukan metode penandatanganan dan klaim.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Menetapkan klaim, dalam hal ini, menetapkan email.
	claims["email"] = email

	// Menghasilkan string token menggunakan kunci rahasia.
	tokenString, err := token.SignedString([]byte(a.Config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Login melakukan proses otentikasi pengguna.
func (a *AuthService) Login(email, password string) (string, error) {
	// Validasi email
	if !utils.IsValidEmail(email) {
		return "", errors.New("format email tidak valid")
	}

	// Validasi kata sandi
	if !utils.IsStrongPassword(password) {
		return "", errors.New("kata sandi tidak cukup kuat")
	}

	// Mengambil pengguna dari basis data berdasarkan email yang diberikan.
	user, err := a.FindUserByEmail(email)
	if err != nil {
		return "", a.Error.HandleError(err)
	}

	// Memeriksa apakah pengguna ada
	if user == nil {
		// Jika pengguna tidak ditemukan, kembalikan kesalahan.
		return "", errors.New("email atau kata sandi tidak valid")
	}

	// Memeriksa apakah kata sandi di-hash
	if strings.HasPrefix(user.Password, "a$") {
		// Kata sandi di-hash, membandingkan kata sandi
		if comparePasswords(password, user.Password) {
			// Menghasilkan token JWT untuk pengguna yang diotentikasi.
			if user.JWTToken == "" {
				token, err := a.GenerateJWTToken(email)
				if err != nil {
					return "", a.Error.HandleError(err)
				}

				// Memperbarui token pengguna di basis data
				user.JWTToken = token
				result := a.DB.Save(&user)
				if result.Error != nil {
					return "", a.Error.HandleError(result.Error)
				}
			}
			return user.JWTToken, nil
		}
	} else {
		// Kata sandi dalam teks biasa, meng-hash dan membandingkan
		if comparePasswords(password, user.Password) {
			// Menghasilkan token JWT untuk pengguna yang diotentikasi.
			if user.JWTToken == "" {
				token, err := a.GenerateJWTToken(email)
				if err != nil {
					return "", a.Error.HandleError(err)
				}

				// Memperbarui token pengguna di basis data
				user.JWTToken = token
				result := a.DB.Save(&user)
				if result.Error != nil {
					return "", a.Error.HandleError(result.Error)
				}

				// Meng-hash kata sandi teks biasa sebelum menyimpannya kembali ke basis data
				hashedPassword, err := hashPassword(password)
				if err != nil {
					return "", a.Error.HandleError(err)
				}
				user.Password = hashedPassword
				result = a.DB.Save(&user)
				if result.Error != nil {
					return "", a.Error.HandleError(result.Error)
				}
			}
			return user.JWTToken, nil
		}
	}

	// Jika pengguna tidak ditemukan atau kata sandi tidak benar, kembalikan kesalahan.
	return "", errors.New("email atau kata sandi tidak valid")
}

// FindUserByEmail mengambil pengguna dari basis data berdasarkan email.
func (a *AuthService) FindUserByEmail(email string) (*models.User, error) {
	// Membuat instance model User
	user := &models.User{}

	// Melakukan kueri basis data untuk mencari pengguna berdasarkan email
	result := a.DB.Where("email =?", email).First(user)

	// Memeriksa kesalahan dalam hasil kueri
	if result.Error != nil {
		// Memeriksa apakah kesalahan disebabkan oleh pengguna tidak ditemukan
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("pengguna tidak ditemukan")
		}

		// Mengembalikan kesalahan lain yang terkait dengan basis data
		return nil, result.Error
	}

	// Mengembalikan pengguna jika ditemukan
	return user, nil
}

// FindAdminCredentials mengambil informasi akun admin dari basis data.
func (a *AuthService) FindAdminCredentials() (*models.User, error) {
	adminUser := &models.User{}
	result := a.DB.Where("role =?", "admin").First(adminUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return adminUser, nil
}

// VerifyToken memverifikasi token JWT.
func (a *AuthService) VerifyToken(tokenString string) (*jwt.Token, error) {
	// Menganalisis token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memeriksa metode penandatanganan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode penandatanganan tidak terduga: %v", token.Header["alg"])
		}

		// Mengembalikan kunci rahasia untuk verifikasi
		return []byte(a.Config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// IsTokenExpired memeriksa apakah token JWT sudah kedaluwarsa.
func (a *AuthService) IsTokenExpired(token *jwt.Token) bool {
	// Mendapatkan waktu kedaluwarsa dari klaim token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return true // Klaim token tidak valid
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

	// Memeriksa apakah token sudah kedaluwarsa
	return expirationTime.Before(time.Now())
}

// comparePasswords membandingkan kata sandi yang diberikan dengan kata sandi yang di-hash dari basis data.
func comparePasswords(providedPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}

// hashPassword meng-hash kata sandi teks biasa menggunakan bcrypt.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// HandleError menangani dan mengembalikan kesalahan dengan respons yang sesuai.
func (a *AuthService) HandleError(err error) error {
	// Merekam kesalahan berdasarkan tingkat log (Anda mungkin memilih untuk merekam dengan cara yang berbeda di sini)
	switch a.Error.LogLevel {
	case "info":
		a.Error.Logger.Println("INFO:", err.Error())
	case "warn":
		a.Error.Logger.Println("WARNING:", err.Error())
	case "error":
		a.Error.Logger.Println("ERROR:", err.Error())
	default:
		// Default untuk mencatat sebagai kesalahan
		a.Error.Logger.Println("ERROR:", err.Error())
	}

	// Mengimplementasikan logika untuk mengembalikan respons yang sesuai.
	var response utils.ErrorResponse

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Message = "Data tidak ditemukan"
	} else if strings.Contains(err.Error(), "unique constraint") {
		response.Message = "Entri duplikat, harap berikan nilai yang unik"
	} else if strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
		response.Message = "Kata sandi tidak valid"
	} else if strings.Contains(err.Error(), "unexpected signing method") {
		response.Message = "Token tidak valid"
	} else if strings.Contains(err.Error(), "Token is expired") {
		response.Message = "Token telah kedaluwarsa"
	} else {
		response.Message = "Kesalahan server internal"
	}

	// Mengembalikan ErrorResponse untuk dikirim sebagai respons JSON
	return &response
}
