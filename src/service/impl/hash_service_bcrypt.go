package impl

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashBcryptService struct {
}

func NewHashBcryptService() *HashBcryptService {
	return &HashBcryptService{}
}

func (h *HashBcryptService) HashPassword(password string, hashSalt string) (string, error) {
	salt, err := strconv.Atoi(hashSalt)

	if err != nil {
		salt = 10
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func (h *HashBcryptService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
