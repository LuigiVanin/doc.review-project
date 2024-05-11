package entity

import "time"

type Homework struct {
	Id           string     `json:"id" db:"id"`
	Title        string     `json:"title" db:"title"`
	Description  string     `json:"description" db:"description"`
	Owner_id     string     `json:"owner_id" db:"owner_id"`
	Deleted      bool       `json:"deleted" db:"deleted"`
	DeadlineDate *time.Time `json:"deadline_date" db:"deadline_date"`
	Code         string     `json:"code" db:"code"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
}
