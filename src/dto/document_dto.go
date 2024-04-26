package dto

type CreateDocumentDto struct {
	Title   string `json:"title" validate:"required,max=255"`
	Content string `json:"content" validate:"required"`
}

type PatchDocumentDto struct {
	Id      string
	Title   *string `json:"title" validate:"omitempty,max=255"`
	Content *string `json:"content" validate:"omitempty"`
}
