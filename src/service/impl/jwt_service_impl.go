package impl

import (
	"doc-review/src/configuration"
	"doc-review/src/exceptions/errors"
	"doc-review/src/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceImpl struct {
	config configuration.Config
}

func NewJwtServiceImpl(config configuration.Config) *JwtServiceImpl {
	return &JwtServiceImpl{config}
}

func (js *JwtServiceImpl) GenerateToken(payload service.JwtPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, service.JwtPayload{
		UserId: payload.UserId,
		Time:   payload.Time,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(payload.Time+3600, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(payload.Time, 0)),
		},
	})

	tokenString, err := token.SignedString([]byte(js.config.Get("JWT_SECRET")))

	if err != nil {
		return "", errors.NewUnauthorizedError("Failed to generate token")
	}

	return tokenString, nil
}

func (js *JwtServiceImpl) VerifyToken(token string) (*service.JwtPayload, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &service.JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.config.Get("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*service.JwtPayload); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.NewUnauthorizedError("Token is not valid")
}
