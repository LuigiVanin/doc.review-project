package dto

// TODO: add updated_at and deadline_date columns to the homework table
type CreateHomeworkDto struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	// OwnerId     string `json:"owner_id" db:"owner_id"`
}
