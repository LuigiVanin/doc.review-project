package configuration

import (
	"doc-review/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: middleware.ErrorHandlerMiddleware,
	}
}
