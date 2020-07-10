package models

import (
	"github.com/go-pg/pg/v9"
)

var dbConnect *pg.DB

//CreateTables Create tables 
func CreateTables(db *pg.DB) error{
	CreateRoleTable(db)
	CreateUserTable(db)

	return nil
}

//InitiateDB initiates a db
func InitiateDB(db *pg.DB){
	dbConnect = db
}
