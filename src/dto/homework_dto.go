package dto

//  TODO: add updated_at and deadline columns to the homework table

type CreateHomeworkDto struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"content" validate:"required"`
}
