package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"
	Enum "doc-review/src/entity/enum"
	"doc-review/src/exceptions/errors"
	"doc-review/src/repository"
	"fmt"
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
	if user.Type == Enum.UserTypeTeacher {
		return entity.Document{}, errors.NewForbiddenError("Only students can create documents")
	}

	return ds.documentRepository.Create(user.Id, document)
}

func (ds *DocumentServiceImpl) ListUserDocuments(userId string) ([]entity.Document, error) {
	return ds.documentRepository.ListUserDocuments(userId)
}

func (ds *DocumentServiceImpl) Update(user dto.ResponseUserDto, document dto.PatchDocumentDto) (entity.Document, error) {
	if document.Id == "" {
		return entity.Document{}, errors.NewBadRequestError("Document ID is required")
	}

	if user.Type == Enum.UserTypeTeacher {
		return entity.Document{}, errors.NewForbiddenError("Only students can update documents")
	}

	foundDocument, err := ds.documentRepository.FindById(document.Id)

	if user.Id != foundDocument.AuthorId || err != nil {
		return entity.Document{}, errors.NewForbiddenError("You can only update your own documents")
	}

	fmt.Println("foundDocument", foundDocument)
	fmt.Println("Document Updated!!")

	return ds.documentRepository.Update(document)
}

func (ds *DocumentServiceImpl) FindById(user dto.ResponseUserDto, documentId string) (entity.Document, error) {
	if documentId == "" {
		return entity.Document{}, errors.NewBadRequestError("Document ID is required")
	}

	foundDocument, err := ds.documentRepository.FindById(documentId)

	if foundDocument.Id == "" || err != nil {
		return entity.Document{}, errors.NewNotFoundError("Document not found")
	}

	if user.Id != foundDocument.AuthorId || err != nil {
		return entity.Document{}, errors.NewForbiddenError("You can only view your own documents")
	}

	return foundDocument, nil
}
