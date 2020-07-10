package models

import (
	"log"
	"time"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
)
// Role type
type  Role struct {
	ID int `josn:"id"`
	Name string `josn:"name"`
	Description string `josn:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// CreateRoleTable creates a user table
func CreateRoleTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Role{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	
	log.Printf("Role table created")
	return nil
}