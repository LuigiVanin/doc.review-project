package impl

import (
	"doc-review/src/dto"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"doc-review/src/service"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(ur repository.UserRepository) service.UserService {
	return &UserServiceImpl{
		userRepository: ur,
	}
}

func (service *UserServiceImpl) FindById(id string) (interface{}, error) {
	panic("implement me")
}

func (service *UserServiceImpl) Create(user dto.CreateUserDto) (dto.GetUserDto, error) {

	_, err := service.userRepository.FindByEmail(user.Email)

	if err == nil {
		return dto.GetUserDto{}, errors.NewConflictError("Email already exists")
	}

	if responseUser, err := service.userRepository.Create(user); err == nil {
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
