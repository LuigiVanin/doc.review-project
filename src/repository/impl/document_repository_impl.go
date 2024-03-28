package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"

	"github.com/jmoiron/sqlx"
)

type DocumentRepositoryImpl struct {
	database *sqlx.DB
}

func NewDocumentRepositoryImpl(database *sqlx.DB) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{database: database}
}

func (repository *DocumentRepositoryImpl) Create(userId string, document dto.CreateDocumentDto) (entity.Document, error) {
	// fmt.Println("Create", document)
	query := "INSERT INTO documents (title, content, author_id) VALUES ($1, $2, $3) RETURNING *"
	var responseDocument entity.Document

	err := repository.database.Get(&responseDocument, query, document.Title, document.Content, userId)

	if err != nil {
		return responseDocument, err
	}

	return responseDocument, nil
}
