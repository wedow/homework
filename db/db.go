package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("postgres", "dbname=homework sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}

// Use allows a different database connection to be specified for future queries
func Use(database *sql.DB) {
	db = database
}
