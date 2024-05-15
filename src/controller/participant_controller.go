package controller

import (
	"doc-review/src/dto"
	"doc-review/src/guard"
	m "doc-review/src/middleware"
	"doc-review/src/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ParticipantController struct {
	participantService service.ParticipantService
	authGuard          guard.Guard
}

func NewParticipantController(ps service.ParticipantService, ag guard.Guard) *ParticipantController {
	return &ParticipantController{
		participantService: ps,
		authGuard:          ag,
	}
}

func (controller *ParticipantController) Add(ctx *fiber.Ctx) error {
	fmt.Println("Adding participant to homework")
	user := ctx.Locals("user").(*dto.ResponseUserDto)
	homeworkCode := ctx.Params("homework_code")

	fmt.Println(homeworkCode)

	homework, err := controller.participantService.AddParticipantToHomework(*user, homeworkCode)

	if err != nil {
		return err
	}

	return ctx.JSON(homework)
}

func (controller *ParticipantController) Register(app *fiber.App) {
	fmt.Println("Registering participant controller")

	app.Post("/homeworks/:homework_code/participants",
		m.JsonValidator[dto.AddParticipantDto](),
		controller.authGuard.Activate,
		controller.Add)
}
