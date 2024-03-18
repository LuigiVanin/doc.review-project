package controller

import (
	"doc-review/src/dto"
	"doc-review/src/middleware"
	"doc-review/src/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (controller *UserController) Register(app *fiber.App) {

	app.Get("/users/:id", controller.FindById)
	app.Post("/users", middleware.JsonValidator[dto.CreateUserDto](), controller.Create)
}

func (controller *UserController) FindById(c *fiber.Ctx) error {
	userId := c.Params("id")

	fmt.Println("FindById", userId)

	return nil
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	user := c.Locals("json").(*dto.CreateUserDto)

	res, err := controller.userService.Create(*user)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(res)
}
