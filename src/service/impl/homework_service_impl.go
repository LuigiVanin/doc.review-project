package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	Enum "doc-review/src/entity/enum"
	"doc-review/src/exceptions/errors"
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
	if user.Type == Enum.UserTypeStudent {
		return entity.Homework{}, errors.NewForbiddenError("Only teachers can create homeworks")
	}

	if createdHomework, err := service.homeworkRepository.Create(user.Id, homework); err != nil {
		return entity.Homework{}, err
	} else {
		return createdHomework, nil
	}
}

func (service *HomeworkServiceImpl) ListUserHomeworks(user dto.ResponseUserDto) ([]entity.Homework, error) {

	if user.Type == Enum.UserTypeTeacher {
		homeworks, err := service.homeworkRepository.ListOwnerHomeworks(user.Id)

		if err != nil {
			return []entity.Homework{}, errors.NewInternalServerError("Error while fetching homeworks")
		}

		return homeworks, nil
	}

	homeworks, err := service.homeworkRepository.ListStudentHomeworks(user.Id)

	if err != nil {
		return []entity.Homework{}, errors.NewInternalServerError("Error while fetching homeworks")
	}

	return homeworks, nil

}
