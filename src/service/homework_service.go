package service

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type HomeworkService interface {
	Create(user dto.ResponseUserDto, homework dto.CreateHomeworkDto) (entity.Homework, error)
	ListUserHomeworks(user dto.ResponseUserDto) ([]entity.Homework, error)
}
