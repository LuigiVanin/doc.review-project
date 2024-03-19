package service

import "doc-review/src/dto"

type UserService interface {
	FindById(id string) (interface{}, error)
	Create(user dto.CreateUserDto) (dto.ResponseUserDto, error)
}
