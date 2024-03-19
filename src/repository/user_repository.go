package repository

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type UserRepository interface {
	FindById(id string) (dto.ResponseUserDto, error)
	FindByEmail(email string) (entity.User, error)
	Create(user dto.CreateUserDto) (entity.User, error)
}
