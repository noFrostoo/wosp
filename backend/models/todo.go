package models

import "github.com/google/uuid"

type Todo struct {
	Id          uuid.UUID `db:"id"`
	User_id     uuid.UUID `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Done        bool      `db:"done"`
	Due_at      string    `db:"due_at"`
}
