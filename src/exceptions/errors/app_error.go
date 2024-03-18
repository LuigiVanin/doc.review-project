package errors

import "github.com/gofiber/fiber/v2"

type FiberError = fiber.Error

type AppError struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Code    string       `json:"code"`
	Fields  []FieldError `json:"fields"`
}

func (e *AppError) FmtFields() []FieldError {
	return e.Fields
}

func (e *AppError) Error() string {
	return e.Message
}
