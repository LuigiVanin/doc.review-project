package repository

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type DocumentRepository interface {
	// FindByEmail(email string) (entity.User, error)
	FindById(id string) (entity.Document, error)
	Create(userId string, document dto.CreateDocumentDto) (entity.Document, error)
	Update(document dto.PatchDocumentDto) (entity.Document, error)
	ListUserDocuments(userId string) ([]entity.Document, error)
}
