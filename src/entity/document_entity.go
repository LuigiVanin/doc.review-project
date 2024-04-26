package entity

import "time"

type Document struct {
	Id                  string    `json:"id" db:"id"`
	Title               string    `json:"title" db:"title"`
	Content             string    `json:"content" db:"content"`
	AuthorId            string    `json:"author_id" db:"author_id"`
	SubmittedHomeworkId *string   `json:"submitted_homework_id" db:"submitted_homework_id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
