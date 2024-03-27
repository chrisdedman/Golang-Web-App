package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

func main() {
	// main connects to the database, performs some operations, and then closes the connection.
	databaseUrl := getEnvVar("DATABASE_URL")

	if len(databaseUrl) == 0 {
		fmt.Fprintf(os.Stderr, "DATABASE_URL environment variable not set\n")
		os.Exit(1)
	}
	// open database
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected!")

	/*
	 * Exemple of how to use the database to create a table,
	 *	insert data and retrieve data

	 *	// test for create table and insert data
	 *	_, err = dbPool.Exec(context.Background(), "create table if not exists users(id serial primary key, username text, email text, password text, created_at timestamp, updated_at timestamp)")
	 *	if err != nil {
	 *		fmt.Fprintf(os.Stderr, "Error creating table: %v\n", err)
	 *		os.Exit(1)
	 *	}

	 *	_, err = dbPool.Exec(context.Background(), "insert into users(username, email, password, created_at, updated_at) values($1, $2, $3, $4, $5)", "Chris", "test@email.com", "1234", "2021-07-01", "2021-07-01")
	 *	if err != nil {
	 *		fmt.Fprintf(os.Stderr, "Error inserting data: %v\n", err)
	 *		os.Exit(1)
	 *	}

	 *	// query database test
	 *	var name string
	 *	err = dbPool.QueryRow(context.Background(), "select username from users").Scan(&name)
	 *	if err != nil {
	 *		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	 *		os.Exit(1)
	 *	}

	 *	fmt.Println(name)
	 */

	// close database
	defer dbPool.Close()
}
