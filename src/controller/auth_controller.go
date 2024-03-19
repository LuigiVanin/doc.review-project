package controller

import (
	"doc-review/src/dto"
	m "doc-review/src/middleware"
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
	user := c.Locals("json").(*dto.SigninDto)

	res, err := controller.userService.Signin(*user)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(res)
}

func (controller *AuthController) Signup(c *fiber.Ctx) error {
	user := c.Locals("json").(*dto.SignupDto)

	res, err := controller.userService.Signup(*user)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(res)
}

func (controller *AuthController) Register(app *fiber.App) {
	app.Post("/auth/signin", m.JsonValidator[dto.SigninDto](), controller.Signin)
	app.Post("/auth/signup", m.JsonValidator[dto.SignupDto](), controller.Signup)
}
