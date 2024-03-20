package service

import "doc-review/src/dto"

type UserService interface {
	FindById(id string) (dto.ResponseUserDto, error)
	Create(user dto.CreateUserDto) (dto.ResponseUserDto, error)
}
