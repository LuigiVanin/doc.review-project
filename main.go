package main

import (
	module "doc-review/src"
	"doc-review/src/configuration"
	"doc-review/src/controller"
	"doc-review/src/guard"
	repository "doc-review/src/repository/impl"
	service "doc-review/src/service/impl"

	_ "github.com/lib/pq"
)

func main() {
	app := module.NewApiApp()

	config := configuration.NewConfig()
	database := configuration.NewDatabase(config)
	defer database.Close()

	userRepository := repository.NewUserRepositoryImpl(database)

	hashService := service.NewHashBcryptService()
	userService := service.NewUserServiceImpl(userRepository)
	jwtService := service.NewJwtServiceImpl(config)
	authService := service.NewAuthServiceImpl(config, hashService, jwtService, userRepository)

	authGuard := guard.NewAuthorizationGuard(userService, jwtService, userRepository)

	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService, authGuard)

	app.Register(authController)
	app.Register(userController)

	app.Start()
}
