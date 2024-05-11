package repository

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type HomeworkRepository interface {
	Create(ownerId string, homework dto.CreateHomeworkDto) (entity.Homework, error)
	ListOwnerHomeworks(userId string) ([]entity.Homework, error)
	ListStudentHomeworks(userId string) ([]entity.Homework, error)
}
