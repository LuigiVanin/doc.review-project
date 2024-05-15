package repository

import "doc-review/src/entity"

type ParticipantRepository interface {
	Create(entity.Participant) (entity.Participant, error)
}
