// app/config/config.go

package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config menyimpan konfigurasi untuk aplikasi.
type Config struct {
	MySQLHost     string
	MySQLPort     string
	MySQLProtocol string
	MySQLUser     string
	MySQLPassword string
	JWTSecret     string
	MySQLName     string
}

// AuthServiceConfigInterface mendefinisikan antarmuka untuk konfigurasi AuthService.
type AuthServiceConfigInterface interface {
	GetAuthServiceConfig() *Config
}

// GetAuthServiceConfig mengembalikan konfigurasi untuk AuthService.
func (c *Config) GetAuthServiceConfig() *Config {
	return &Config{
		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLProtocol: getEnv("MYSQL_PROTOCOL", "tcp"),
		MySQLUser:     getEnv("MYSQL_USER", "root"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "123456789"),
		MySQLName:     getEnv("MYSQL_NAME", "wp"),
		JWTSecret:     getEnv("JWT_SECRET", "dibimbing"),
	}
}

// LoadConfig memuat konfigurasi umum aplikasi.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	return &Config{
		MySQLHost:     getEnv("MYSQL_HOST", "localhost"),
		MySQLPort:     getEnv("MYSQL_PORT", "3306"),
		MySQLProtocol: getEnv("MYSQL_PROTOCOL", "tcp"),
		MySQLUser:     getEnv("MYSQL_USER", "root"),
		MySQLPassword: getEnv("MYSQL_PASSWORD", "123456789"),
		MySQLName:     getEnv("MYSQL_NAME", "wp"),
		JWTSecret:     getEnv("JWT_SECRET", "dibimbing"),
	}, nil
}

// InitializeDBConnection menginisialisasi koneksi database.
func InitializeDBConnection() (*gorm.DB, error) {
	// Menghapus parameter 'config *Config' untuk menghindari siklus impor
	dsn := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?parseTime=true",
		getEnv("MYSQL_USER", "root"),
		getEnv("MYSQL_PASSWORD", "123456789"),
		getEnv("MYSQL_PROTOCOL", "tcp"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_NAME", "wp"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke database: %v", err)
	}

	return db, nil
}

// getEnv mengambil nilai dari variabel lingkungan atau mengembalikan nilai default.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
