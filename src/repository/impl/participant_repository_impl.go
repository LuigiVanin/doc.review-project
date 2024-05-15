package impl

import (
	"doc-review/src/entity"

	"github.com/jmoiron/sqlx"
)

type ParticipantRepositoryImpl struct {
	database *sqlx.DB
}

func NewParticipantRepositoryImpl(database *sqlx.DB) *ParticipantRepositoryImpl {
	return &ParticipantRepositoryImpl{
		database,
	}
}

func (repository *ParticipantRepositoryImpl) Create(participant entity.Participant) (entity.Participant, error) {
	query := "INSERT INTO participants (user_id, homework_id) VALUES ($1, $2) RETURNING *"
	var newParticipant entity.Participant

	err := repository.database.Get(
		&newParticipant,
		query,
		participant.UserId,
		participant.HomeworkId,
	)

	if err != nil {
		return newParticipant, err
	}

	return newParticipant, nil

}
