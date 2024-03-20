package guard

import fiber "github.com/gofiber/fiber/v2"

type Guard interface {
	Activate(ctx *fiber.Ctx) error
}
