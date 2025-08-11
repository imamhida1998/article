package lib

import (
	"article/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectiontoPostgreSQL(config model.DBConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Connected to PostgreSQL database successfully")
	}

	return db, nil
}
