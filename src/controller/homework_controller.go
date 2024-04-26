package controller

import (
	"doc-review/src/dto"
	"doc-review/src/guard"
	m "doc-review/src/middleware"
	"doc-review/src/service"

	"github.com/gofiber/fiber/v2"
)

type HomeworkController struct {
	homeworkService service.HomeworkService
	authGuard       guard.Guard
}

func NewHomeworkController(hs service.HomeworkService, ag guard.Guard) *HomeworkController {
	return &HomeworkController{
		homeworkService: hs,
		authGuard:       ag,
	}
}

func (controller *HomeworkController) Create(ctx *fiber.Ctx) error {
	documentBody := ctx.Locals("json").(*dto.CreateHomeworkDto)
	user := ctx.Locals("user").(*dto.ResponseUserDto)

	homework, err := controller.homeworkService.Create(*user, *documentBody)

	if err != nil {
		return err
	}

	return ctx.Status(201).JSON(homework)

}

func (controller *HomeworkController) Register(app *fiber.App) {

	app.Post("/homeworks",
		m.JsonValidator[dto.CreateDocumentDto](),
		controller.authGuard.Activate,
		controller.Create)
}
