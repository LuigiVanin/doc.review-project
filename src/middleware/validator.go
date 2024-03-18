package middleware

import (
	helpers "doc-review/src/common"
	errors "doc-review/src/exceptions/errors"

	"github.com/gofiber/fiber/v2"
)

func JsonValidator[T interface{}]() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		data, err := helpers.ParseContextBody[T](ctx)

		if err != nil {
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}

		if err := helpers.Validation(data); err != nil {
			if fieldsErr, ok := err.(*errors.FieldErrors); ok {
				return &errors.AppError{
					Status:  fiber.ErrBadRequest.Code,
					Message: fiber.ErrBadRequest.Message,
					Code:    "bad-request",
					Fields:  fieldsErr.Fields,
				}
			}

			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}
		ctx.Locals("json", data)
		return ctx.Next()
	}
}
