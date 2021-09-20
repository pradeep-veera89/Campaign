package models

import "time"

type Lead struct {
	ID         int
	Salutation string
	FirstName  string
	LastName   string
	EMail      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
