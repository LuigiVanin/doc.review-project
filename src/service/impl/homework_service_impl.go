package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	"doc-review/src/repository"
)

type HomeworkServiceImpl struct {
	homeworkRepository repository.HomeworkRepository
}

func NewHomeworkServiceImpl(hr repository.HomeworkRepository) *HomeworkServiceImpl {
	return &HomeworkServiceImpl{
		homeworkRepository: hr,
	}
}

func (service *HomeworkServiceImpl) Create(user dto.ResponseUserDto, homework dto.CreateHomeworkDto) (entity.Homework, error) {
	panic("implement me")
}
