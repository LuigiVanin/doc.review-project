package helpers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ParseContextBody[T interface{}](ctx *fiber.Ctx) (*T, error) {
	var data T
	if err := ctx.BodyParser(&data); err != nil {
		return nil, errors.New("invalid request")
	}
	return &data, nil
}
