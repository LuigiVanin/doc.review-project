package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	helpers "doc-review/src/lib"
	"doc-review/src/repository"
	"fmt"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type HomeworkRepositoryImpl struct {
	database *sqlx.DB
}

func NewHomeworkRepositoryImpl(database *sqlx.DB) *HomeworkRepositoryImpl {
	return &HomeworkRepositoryImpl{database}
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

func (r *HomeworkRepositoryImpl) FindUniqueBy(params repository.HomeworkFindUniqueByParams) (entity.Homework, error) {
	query := "SELECT * FROM homeworks WHERE "
	var homework entity.Homework
	var uniqueIdentifier string

	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.FieldByName(field.Name).String()
		if value != "" {
			dbTag := field.Tag.Get("db")
			uniqueIdentifier = value
			query += fmt.Sprintf("%s = $1", dbTag)
			break
		}
	}

	fmt.Println("TESTE ")

	err := r.database.Get(&homework, query, uniqueIdentifier)

	if err != nil {
		return homework, err
	}

	fmt.Println("HOMEWORK: ", homework)

	return homework, nil
}
