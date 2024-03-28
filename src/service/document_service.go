package service

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
)

type DocumentService interface {
	Create(user dto.ResponseUserDto, document dto.CreateDocumentDto) (entity.Document, error)
}
