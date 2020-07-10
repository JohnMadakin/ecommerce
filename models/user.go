package models

import (
	"log"
	"time"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
)
// User type
type  User struct {
	ID int `josn:"id"`
	Firstname string `josn:"firstname"`
	Lastname string `josn:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RoleID int `pg:"on_delete:RESTRICT, on_update: CASCADE"`
  Role   *Role
}

// CreateUserTable creates a user table
func CreateUserTable(db *pg.DB) error{
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&User{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("User table created")
	return nil
}