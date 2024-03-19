package service

import "github.com/golang-jwt/jwt/v5"

type JwtPayload struct {
	UserId string `json:"userId"`
	Time   int64  `json:"time"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(data JwtPayload) (string, error)
	VerifyToken(token string) (*JwtPayload, error)
}
