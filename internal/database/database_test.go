package main_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func TestMain(t *testing.T) {
	// Set up a test environment
	databaseUrl := getEnvVar("DATABASE_URL")

	// Connect to the database
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}

	// Close the database connection
	dbPool.Close()
}
