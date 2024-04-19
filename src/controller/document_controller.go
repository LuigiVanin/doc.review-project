package controller

import (
	"doc-review/src/dto"
	"doc-review/src/guard"
	m "doc-review/src/middleware"
	"doc-review/src/service"

	"github.com/gofiber/fiber/v2"
)

type DocumentController struct {
	documentService service.DocumentService
	authGuard       guard.Guard
}

func NewDocumentController(ds service.DocumentService, ag guard.Guard) *DocumentController {
	return &DocumentController{
		documentService: ds,
		authGuard:       ag,
	}
}

func (controller *DocumentController) Create(c *fiber.Ctx) error {

	documentBody := c.Locals("json").(*dto.CreateDocumentDto)
	user := c.Locals("user").(*dto.ResponseUserDto)

	document, err := controller.documentService.Create(*user, *documentBody)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(document)
}

func (controller *DocumentController) ListUserDocuments(c *fiber.Ctx) error {
	user := c.Locals("user").(*dto.ResponseUserDto)
	document, err := controller.documentService.ListUserDocuments(user.Id)

	if err != nil {
		return err
	}

	return c.JSON(document)
}

func (controller *DocumentController) Register(app *fiber.App) {
	app.Post(
		"/documents",
		m.JsonValidator[dto.CreateDocumentDto](),
		controller.authGuard.Activate,
		controller.Create,
	)
	app.Get(
		"/documents",
		controller.authGuard.Activate,
		controller.ListUserDocuments,
	)
}
