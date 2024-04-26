package guard

import (
	Enum "doc-review/src/entity/enum"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"

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

	// TODO: Add Bearer check on the start of the Authorization Header

	payload, err := ag.jwtService.VerifyToken(header.Authorization)

	if err != nil {
		return errors.NewUnauthorizedError("Bad Formatted token")
	}

	user, err := ag.userService.FindById(payload.UserId)

	// TODO: Add jwt time validation

	if err != nil {
		return errors.NewUnauthorizedError("User not found")
	}

	ctx.Locals(Enum.LocalsUser, &user)

	return ctx.Next()
}
