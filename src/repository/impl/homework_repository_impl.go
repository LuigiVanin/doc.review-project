package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	helpers "doc-review/src/lib"

	"github.com/jmoiron/sqlx"
)

type HomeworkRepositoryImpl struct {
	database *sqlx.DB
}

func NewHomeworkRepositoryImpl(database *sqlx.DB) *HomeworkRepositoryImpl {
	return &HomeworkRepositoryImpl{database: database}
}

func (r *HomeworkRepositoryImpl) Create(ownerId string, homework dto.CreateHomeworkDto) (entity.Homework, error) {
	query := "INSERT INTO homeworks (title, description,  owner_id, code) VALUES ($1, $2, $3, $4) RETURNING *"
	var newHomework entity.Homework

	code_id := helpers.GenerateID(7)

	err := r.database.Get(
		&newHomework,
		query,
		homework.Title,
		homework.Description,
		ownerId,
		code_id,
	)

	if err != nil {
		return newHomework, err
	}

	return newHomework, nil
}

func (r *HomeworkRepositoryImpl) ListOwnerHomeworks(userId string) ([]entity.Homework, error) {
	query := "SELECT * FROM homeworks WHERE owner_id = $1"
	var homeworks []entity.Homework

	err := r.database.Select(&homeworks, query, userId)

	if err != nil {
		return homeworks, err
	}

	return homeworks, nil
}

func (r *HomeworkRepositoryImpl) ListStudentHomeworks(userId string) ([]entity.Homework, error) {
	query := "SELECT * FROM homeworks WHERE id IN (SELECT homework_id FROM homework_students WHERE student_id = $1)"
	var homeworks []entity.Homework

	err := r.database.Select(&homeworks, query, userId)

	if err != nil {
		return homeworks, err
	}

	return homeworks, nil
}
