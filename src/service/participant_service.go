package service

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type ParticipantService interface {
	AddParticipantToHomework(user dto.ResponseUserDto, homework_id string) (entity.Participant, error)
	// ListHomeworkParticipants(homeworkId string) ([]entity.Participant, error)
}
