package guard

import (
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationGuard struct {
	userRepository repository.UserRepository
	userService    service.UserService
	jwtService     service.JwtService
}

func NewAuthorizationGuard(
	us service.UserService,
	js service.JwtService,
	ur repository.UserRepository,
) *AuthorizationGuard {

	return &AuthorizationGuard{
		userService:    us,
		jwtService:     js,
		userRepository: ur,
	}
}

func (ag *AuthorizationGuard) Activate(ctx *fiber.Ctx) error {
	type AuthHeader struct {
		Authorization string `reqHeader:"Authorization"`
	}

	header := AuthHeader{}

	if err := ctx.ReqHeaderParser(&header); err != nil {
		return err
	}

	payload, err := ag.jwtService.VerifyToken(header.Authorization)

	if err != nil {
		return errors.NewUnauthorizedError("Bad Formatted token")
	}
	fmt.Println(payload.UserId)

	user, err := ag.userService.FindById(payload.UserId)

	if err != nil {
		return errors.NewUnauthorizedError("User not found")
	}

	ctx.Locals("user", user)

	return ctx.Next()
}
