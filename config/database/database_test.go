package database_test

import (
	"os"
	"testing"

	"github.com/golang-web-app/config/database"
	"github.com/golang-web-app/internal/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func TestInitDB(t *testing.T) {
	err := godotenv.Load("../../.env")
	assert.NoError(t, err)

	config := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.InitDB(config)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.User{Username: "testDev", Email: "testdev@test.dev", Password: "Password123!@#"})
	assert.NoError(t, err)
}
