package service

import "doc-review/src/dto"

type AuthService interface {
	Signin(creadential dto.SigninDto) (dto.GetUserDto, error)
}
