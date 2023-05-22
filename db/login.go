package db

import (
	"github.com/John-Santa/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.Usuario, bool) {
	user, found, _ := ExistUser(email)
	if !found {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
