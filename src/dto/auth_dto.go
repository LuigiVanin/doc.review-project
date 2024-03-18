package dto

type SigninDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type SignupDto struct {
	CreateUserDto
	PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
}
