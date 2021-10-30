package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string,error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password),12)

	return string(bytes),err;
}


func ComparePassword(hashPassword string, password string) bool {

	isValid:= bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))

	return isValid == nil;
}