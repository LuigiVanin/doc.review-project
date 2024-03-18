package service

type HashService interface {
	HashPassword(password string, hashSalt string) (string, error)
	ComparePassword(hashed string, text string) bool
}
