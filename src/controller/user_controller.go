package controller

import (
	"doc-review/src/dto"
	"doc-review/src/exceptions/errors"
	"doc-review/src/guard"
	"doc-review/src/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
	authGuard   guard.Guard
}

func NewUserController(us service.UserService, ag guard.Guard) *UserController {
	return &UserController{
		userService: us,
		authGuard:   ag,
	}
}

func (controller *UserController) FindById(c *fiber.Ctx) error {
	userId := c.Params("id")
	authUser := c.Locals("user").(*dto.ResponseUserDto)

	if userId == authUser.Id {
		return c.Status(200).JSON(authUser)
	}

	return errors.NewUnauthorizedError("Unauthorized access to user data")
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	user := c.Locals("json").(*dto.CreateUserDto)

	res, err := controller.userService.Create(*user)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(res)
}

func (controller *UserController) Register(app *fiber.App) {

	app.Get("/users/:id", controller.authGuard.Activate, controller.FindById)
}
