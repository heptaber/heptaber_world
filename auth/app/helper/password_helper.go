package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashed)
}

func VerifyPassword(hashedUserPassword string, providedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedUserPassword), []byte(providedPassword))
}
