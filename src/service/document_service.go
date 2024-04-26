package service

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type DocumentService interface {
	Create(user dto.ResponseUserDto, document dto.CreateDocumentDto) (entity.Document, error)
	Update(user dto.ResponseUserDto, document dto.PatchDocumentDto) (entity.Document, error)
	FindById(user dto.ResponseUserDto, documentId string) (entity.Document, error)
	ListUserDocuments(userId string) ([]entity.Document, error)
}
