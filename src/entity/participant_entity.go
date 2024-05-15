package entity

import "time"

type Participant struct {
	Id         string    `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	HomeworkId string    `json:"homework_id" db:"homework_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
