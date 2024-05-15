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
	api := module.NewApiApp()

	config := configuration.NewConfig()
	database := configuration.NewDatabase(config)
	defer database.Close()

	userRepository := repository.NewUserRepositoryImpl(database)
	documentRepository := repository.NewDocumentRepositoryImpl(database)
	homeworkRepository := repository.NewHomeworkRepositoryImpl(database)
	participantRepository := repository.NewParticipantRepositoryImpl(database)

	hashService := service.NewHashBcryptService()
	userService := service.NewUserServiceImpl(userRepository)
	jwtService := service.NewJwtServiceImpl(config)
	authService := service.NewAuthServiceImpl(config, hashService, jwtService, userRepository)
	documentService := service.NewDocumentServiceImpl(documentRepository)
	homeworkService := service.NewHomeworkServiceImpl(homeworkRepository)
	participantService := service.NewParticipantServiceImpl(participantRepository, homeworkRepository)

	authGuard := guard.NewAuthorizationGuard(userService, jwtService, userRepository)

	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService, authGuard)
	documentController := controller.NewDocumentController(documentService, authGuard)
	homeworkController := controller.NewHomeworkController(homeworkService, authGuard)
	participantController := controller.NewParticipantController(participantService, authGuard)

	api.Register(authController)
	api.Register(userController)
	api.Register(documentController)
	api.Register(participantController)
	api.Register(homeworkController)

	api.Start()
}
