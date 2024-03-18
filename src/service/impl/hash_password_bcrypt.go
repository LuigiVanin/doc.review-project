package impl

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashPasswordBcrypt struct {
}

func NewHashPasswordBcrypt() *HashPasswordBcrypt {
	return &HashPasswordBcrypt{}
}

func (h *HashPasswordBcrypt) HashPassword(password string, hashSalt string) (string, error) {
	salt, err := strconv.Atoi(hashSalt)

	if err != nil {
		salt = 10
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func (h *HashPasswordBcrypt) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
