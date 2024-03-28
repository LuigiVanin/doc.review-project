package middleware

import (
	errors "doc-review/src/exceptions/errors"
	helpers "doc-review/src/lib"

	"github.com/gofiber/fiber/v2"
)

func ParseContextBody[T interface{}](ctx *fiber.Ctx) (*T, *errors.AppError) {
	var data T
	if err := ctx.BodyParser(&data); err != nil {
		return nil, errors.NewBadRequestError("Invalid request body")
	}
	return &data, nil
}

func JsonValidator[T interface{}]() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		data, err := ParseContextBody[T](ctx)

		if err != nil {
			return err
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
