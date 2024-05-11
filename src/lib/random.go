package helpers

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID(length int) string {
	id := make([]byte, length)

	for i := range id {
		id[i] = charset[rand.Intn(len(charset))]
	}
	return string(id)
}
