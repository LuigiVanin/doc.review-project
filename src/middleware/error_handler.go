package middleware

import (
	"doc-review/src/exceptions/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandlerMiddleware(ctx *fiber.Ctx, err error) error {

	if appErr, ok := err.(*errors.AppError); ok {
		return ctx.Status(appErr.Status).JSON(appErr)
	}

	if fiberErr, ok := err.(*fiber.Error); ok {
		return ctx.Status(fiberErr.Code).JSON(fiber.Map{
			"status":  fiberErr.Code,
			"message": fiberErr.Message,
		})
	}

	return ctx.Status(500).JSON(errors.NewInternalServerError(err.Error()))
}
