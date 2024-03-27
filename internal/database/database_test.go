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
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func TestMain(t *testing.T) {
	// Set up a test environment
	// TO-DO: Fix the path to the .env file not recognized by godotenv
	databaseUrl := getEnvVar("DATABASE_URL")

	// Perform assertions or additional tests if needed
	// For example, you can check if the table and data were created successfully
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}

	var count int
	err = dbPool.QueryRow(context.Background(), "select count(*) from users").Scan(&count)
	if err != nil {
		t.Fatalf("QueryRow failed: %v", err)
	}

	expectedCount := 1
	if count != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, count)
	}

	// Clean up the test environment if needed
	_, err = dbPool.Exec(context.Background(), "drop table if exists users")
	if err != nil {
		t.Fatalf("Error dropping table: %v", err)
	}

	// Close the database connection
	dbPool.Close()
}
