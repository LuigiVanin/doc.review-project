package impl

import (
	"doc-review/src/configuration"
	"doc-review/src/dto"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
	hashService    service.HashService
	config         configuration.Config
}

func NewAuthServiceImpl(config configuration.Config, hs service.HashService, ur repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: ur,
		hashService:    hs,
		config:         config,
	}
}

func (service *AuthServiceImpl) Signin(creadential dto.SigninDto) (dto.GetUserDto, error) {
	user, err := service.userRepository.FindByEmail(creadential.Email)

	if err != nil {
		return dto.GetUserDto{}, errors.NewNotFoundError("User not found")
	}

	if service.hashService.ComparePassword(creadential.Password, user.Password) {
		return dto.GetUserDto{
			Id:        user.Id,
			Type:      user.Type,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}, nil
	} else {
		return dto.GetUserDto{}, errors.NewUnauthorizedError("Invalid password")
	}
}

func (service *AuthServiceImpl) Signup(user dto.SignupDto) (dto.GetUserDto, error) {
	_, err := service.userRepository.FindByEmail(user.Email)

	if err == nil {
		return dto.GetUserDto{}, errors.NewConflictError("Email already exists")
	}

	hashedPassword, err := service.hashService.HashPassword(user.Password, service.config.Get("HASH_SALT"))

	if err != nil {
		return dto.GetUserDto{}, errors.NewInternalServerError("Failed to hash password")
	}

	user.Password = hashedPassword

	if responseUser, err := service.userRepository.Create(user.CreateUserDto); err == nil {
		return dto.GetUserDto{
			Id:        responseUser.Id,
			Type:      responseUser.Type,
			Name:      responseUser.Name,
			Email:     responseUser.Email,
			CreatedAt: responseUser.CreatedAt,
			UpdatedAt: responseUser.UpdatedAt,
		}, nil
	} else {
		return dto.GetUserDto{}, err
	}
}
