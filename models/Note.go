package models

import "time"

type Note struct {
	NoteID      int
	Title       string
	Description string
	CreateAt    time.Time
	UpdatedAt   time.Time
}
