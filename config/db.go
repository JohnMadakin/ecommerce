package config

import (
"log"
"os"
"github.com/go-pg/pg/v9"
)

// Connect Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Addr: os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_NAME"),
	}
	var db *pg.DB = pg.Connect(opts)

	if db == nil {
	log.Printf("Failed to connect")
	os.Exit(100)
	}

	log.Printf("Connected to db")
	return db
}