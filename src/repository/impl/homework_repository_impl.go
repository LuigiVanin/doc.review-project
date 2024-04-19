package impl

import "github.com/jmoiron/sqlx"

type HomeworkRepositoryImpl struct {
	database *sqlx.DB
}

func NewHomeworkRepositoryImpl(database *sqlx.DB) *HomeworkRepositoryImpl {
	return &HomeworkRepositoryImpl{database: database}
}
