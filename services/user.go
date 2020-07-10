package services

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
)

type  User struct {
	ID int `josn:"id"`
	Firstname string `josn:"firstname"`
	Lastname string `josn:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(db *pg.DB) User{
}