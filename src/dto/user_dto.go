package dto

import "time"

type CreateUserDto struct {
	Name     string `json:"name" validate:"required,min=3,max=100" db:"name"`
	Email    string `json:"email" validate:"required,email" db:"email"`
	Password string `json:"password" validate:"required,min=6,max=100" db:"password"`
	Type     string `json:"type" validate:"required,oneof=TEACHER STUDENT" db:"type"`
}

type ResponseUserDto struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type AuthUserDto = ResponseUserDto
