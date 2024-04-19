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

func (repository *DocumentRepositoryImpl) ListUserDocuments(userId string) ([]entity.Document, error) {
	query := "SELECT * FROM documents WHERE author_id = $1"
	var documents []entity.Document

	err := repository.database.Select(&documents, query, userId)

	if err != nil {
		return documents, err
	}

	return documents, nil
}

func (repository *DocumentRepositoryImpl) ListHomeworkDocuments(homeworkId string) ([]entity.Document, error) {
	query := "SELECT * FROM documents WHERE homeword_id = $1"
	var documents []entity.Document

	err := repository.database.Select(&documents, query, homeworkId)

	if err != nil {
		return documents, err
	}

	return documents, nil
}

func (repository *DocumentRepositoryImpl) UpdateDocument(documentId string, document dto.CreateDocumentDto) (entity.Document, error) {
	query := "UPDATE documents SET title = $1, content = $2 WHERE id = $3 RETURNING *"
	var responseDocument entity.Document

	err := repository.database.Get(&responseDocument, query, document.Title, document.Content, documentId)

	if err != nil {
		return responseDocument, err
	}

	return responseDocument, nil
}
