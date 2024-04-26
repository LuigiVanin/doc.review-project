package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DocumentRepositoryImpl struct {
	database *sqlx.DB
}

func NewDocumentRepositoryImpl(database *sqlx.DB) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{database: database}
}

func (repository *DocumentRepositoryImpl) Create(userId string, document dto.CreateDocumentDto) (entity.Document, error) {
	query := "INSERT INTO documents (title, content, author_id) VALUES ($1, $2, $3) RETURNING *"
	var responseDocument entity.Document

	err := repository.database.Get(&responseDocument, query, document.Title, document.Content, userId)

	if err != nil {
		return responseDocument, err
	}

	return responseDocument, nil
}

func (repository *DocumentRepositoryImpl) Update(document dto.PatchDocumentDto) (entity.Document, error) {
	if document.Id == "" || document.Title == nil && document.Content == nil {
		return entity.Document{}, errors.New("document ID, Title or Content is required")
	}

	query := "UPDATE documents SET "
	var responseDocument entity.Document

	queryParams := map[string]interface{}{
		"id": document.Id,
	}

	if document.Title != nil {
		query += "title = :title, "
		queryParams["title"] = *document.Title
	}

	if document.Content != nil {
		query += "content = :content, "
		queryParams["content"] = *document.Content
	}

	query += "updated_at = now() WHERE id = :id"

	_, err := repository.database.NamedExec(query, queryParams)

	if err != nil {
		return responseDocument, err
	}

	responseDocument, err = repository.FindById(document.Id)

	fmt.Println("ERROR: ", err)

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
	query := "SELECT * FROM documents WHERE homework_id = $1"
	var documents []entity.Document

	err := repository.database.Select(&documents, query, homeworkId)

	if err != nil {
		return documents, err
	}

	return documents, nil
}

func (repository *DocumentRepositoryImpl) FindById(id string) (entity.Document, error) {
	query := "SELECT * FROM documents WHERE id = $1"
	var document entity.Document

	err := repository.database.Get(&document, query, id)

	return document, err
}
