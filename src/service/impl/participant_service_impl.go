package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	Enum "doc-review/src/entity/enum"
	"doc-review/src/exceptions/errors"
	r "doc-review/src/repository"
	"fmt"
)

type ParticipantServiceImpl struct {
	participantRepository r.ParticipantRepository
	homeworkRepository    r.HomeworkRepository
}

func NewParticipantServiceImpl(pr r.ParticipantRepository, hp r.HomeworkRepository) *ParticipantServiceImpl {
	return &ParticipantServiceImpl{
		homeworkRepository:    hp,
		participantRepository: pr,
	}
}

type FindUniqueParams = r.HomeworkFindUniqueByParams

func (service *ParticipantServiceImpl) AddParticipantToHomework(user dto.AuthUserDto, homework_code string) (entity.Participant, error) {
	fmt.Println(homework_code)

	// TODO: use user id to define the user to be added

	homework, err := service.homeworkRepository.FindUniqueBy(FindUniqueParams{Code: homework_code})

	if user.Type == Enum.UserTypeTeacher {
		return entity.Participant{}, errors.NewForbiddenError("Only students can be added to homework")
	}

	if err != nil {
		return entity.Participant{}, err
	}

	if homework.Id == "" {
		return entity.Participant{}, errors.NewNotFoundError("Homework not found")
	}

	return service.participantRepository.Create(entity.Participant{UserId: user.Id, HomeworkId: homework.Id})
}
