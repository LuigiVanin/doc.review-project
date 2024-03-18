package module

import (
	"doc-review/src/configuration"
	"doc-review/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Controller interface {
	Register(app *fiber.App)
}

type ApiApp struct {
	app *fiber.App
}

func NewApiApp() *ApiApp {
	app := fiber.New(configuration.NewFiberConfiguration())

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(middleware.Json)

	return &ApiApp{
		app: app,
	}
}

func (ApiApp *ApiApp) Register(controller Controller) {

	controller.Register(ApiApp.app)
}

func (ApiApp *ApiApp) Use(args ...interface{}) fiber.Router {
	return ApiApp.app.Use(args...)
}

func (ApiApp *ApiApp) Start() {
	ApiApp.app.Listen(":3000")
}
