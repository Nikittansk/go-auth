package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB      *sql.DB
	dbError error
)

func Connect(connectionString string) {
	DB, dbError = sql.Open("postgres", connectionString)
	if dbError != nil {
		log.Fatal(dbError)
		return
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connected to PostgreSQL!")
}
