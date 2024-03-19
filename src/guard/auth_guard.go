package guard

import (
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationGuard struct {
	userRepository repository.UserRepository
	jwtService     service.JwtService
}

func NewAuthorizationGuard(
	ur repository.UserRepository,
	js service.JwtService,
) *AuthorizationGuard {
	return &AuthorizationGuard{
		userRepository: ur,
		jwtService:     js,
	}
}

func (ag *AuthorizationGuard) activate(ctx *fiber.Ctx) error {
	type AuthHeader struct {
		Authorization string `reqHeader:"Authorization"`
	}

	header := AuthHeader{}

	if err := ctx.ReqHeaderParser(&header); err != nil {
		return err
	}

	payload, err := ag.jwtService.VerifyToken(header.Authorization)

	if err != nil {
		return err
	}

	user, err := ag.userRepository.FindById(payload.UserId)

	if err != nil {
		return errors.NewUnauthorizedError("User not found")
	}

	ctx.Locals("user", user)

	return ctx.Next()
}
