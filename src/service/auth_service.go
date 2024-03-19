package service

import "doc-review/src/dto"

type SigninResponse struct {
	User  dto.ResponseUserDto
	Token string
}

type AuthService interface {
	Signin(creadential dto.SigninDto) (SigninResponse, error)
	Signup(user dto.SignupDto) (dto.ResponseUserDto, error)
}
