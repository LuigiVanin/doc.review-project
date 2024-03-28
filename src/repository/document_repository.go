package repository

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type DocumentRepository interface {
	// FindById(id string) (entity.User, error)
	// FindByEmail(email string) (entity.User, error)
	Create(userId string, document dto.CreateDocumentDto) (entity.Document, error)
}
