package utils

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

//HashPassword hashes a string 
func HashPassword(password string)string{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte (password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
