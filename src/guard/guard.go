package guard

import fiber "github.com/gofiber/fiber/v2"

type Guard interface {
	activate(ctx *fiber.Ctx) error
}
