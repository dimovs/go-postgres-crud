package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func InitDB() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "dimovs")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "go_crud")

	fmt.Printf("DB Connection params: host=%s, port=%s, user=%s, dbname=%s\n", host, port, user, dbname)

	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", 
		user, password, host, port, dbname)
	fmt.Printf("Connection string: %s\n", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR(100) NOT NULL,
		    email VARCHAR(100) UNIQUE NOT NULL,
		    created_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database")
	return db, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
