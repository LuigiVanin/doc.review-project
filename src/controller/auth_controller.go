package controller

import (
	"doc-review/src/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userService service.AuthService
}

func NewAuthController(us service.AuthService) *AuthController {
	return &AuthController{
		userService: us,
	}
}

func (controller *AuthController) Signin(c *fiber.Ctx) error {
	return nil
}

func (controller *AuthController) Signup(c *fiber.Ctx) error {
	return nil
}

func (controller *AuthController) Register(app *fiber.App) {
	app.Post("/auth/signin", controller.Signin)
	app.Post("/auth/signup", controller.Signup)
}
