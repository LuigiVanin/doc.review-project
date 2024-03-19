package impl

import (
	"doc-review/src/configuration"
	"doc-review/src/dto"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"
	"time"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
	hashService    service.HashService
	config         configuration.Config
	jwtService     service.JwtService
}

func NewAuthServiceImpl(
	config configuration.Config,
	hs service.HashService,
	js service.JwtService,
	ur repository.UserRepository,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: ur,
		hashService:    hs,
		jwtService:     js,
		config:         config,
	}
}

func (as *AuthServiceImpl) Signin(creadential dto.SigninDto) (service.SigninResponse, error) {
	user, err := as.userRepository.FindByEmail(creadential.Email)

	if err != nil {
		return service.SigninResponse{}, errors.NewNotFoundError("User not found")
	}

	if as.hashService.ComparePassword(user.Password, creadential.Password) {
		payload := service.JwtPayload{
			UserId: user.Id,
			Time:   time.Now().Unix(),
		}

		token, err := as.jwtService.GenerateToken(payload)

		if err != nil {
			return service.SigninResponse{}, err
		}

		return service.SigninResponse{
			User: dto.ResponseUserDto{
				Id:        user.Id,
				Type:      user.Type,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
			Token: token,
		}, nil

	}
	return service.SigninResponse{}, errors.NewUnauthorizedError("Invalid password")

}

func (as *AuthServiceImpl) Signup(user dto.SignupDto) (dto.ResponseUserDto, error) {
	_, err := as.userRepository.FindByEmail(user.Email)

	if err == nil {
		return dto.ResponseUserDto{}, errors.NewConflictError("Email already exists")
	}

	hashedPassword, err := as.hashService.HashPassword(user.Password, as.config.Get("HASH_SALT"))

	if err != nil {
		return dto.ResponseUserDto{}, errors.NewInternalServerError("Failed to hash password")
	}

	user.Password = hashedPassword

	if responseUser, err := as.userRepository.Create(user.CreateUserDto); err == nil {
		return dto.ResponseUserDto{
			Id:        responseUser.Id,
			Type:      responseUser.Type,
			Name:      responseUser.Name,
			Email:     responseUser.Email,
			CreatedAt: responseUser.CreatedAt,
			UpdatedAt: responseUser.UpdatedAt,
		}, nil
	} else {
		return dto.ResponseUserDto{}, err
	}
}
