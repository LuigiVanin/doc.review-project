package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
)

type DocumentServiceImpl struct {
	documentRepository repository.DocumentRepository
}

func NewDocumentServiceImpl(dr repository.DocumentRepository) *DocumentServiceImpl {
	return &DocumentServiceImpl{
		documentRepository: dr,
	}
}

func (ds *DocumentServiceImpl) Create(user dto.ResponseUserDto, document dto.CreateDocumentDto) (entity.Document, error) {
	if user.Type == "Professor" {
		return entity.Document{}, errors.NewForbiddenError("Only students can create documents")
	}

	return ds.documentRepository.Create(user.Id, document)
}

func (ds *DocumentServiceImpl) ListUserDocuments(userId string) ([]entity.Document, error) {
	return ds.documentRepository.ListUserDocuments(userId)
}
