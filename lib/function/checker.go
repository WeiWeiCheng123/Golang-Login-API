package function

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckUserIsAccept(username string) error {
	if _, err := strconv.Atoi(username); err == nil {
		return errors.New("must have at least one character and six words")
	}
	if len(username) < 6 {
		return errors.New("must have at least one character and six words")
	}
	return nil
}

func Compare(hash_password string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
}