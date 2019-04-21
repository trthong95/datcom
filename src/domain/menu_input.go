package domain

import "time"

// MenuInput ..
type MenuInput struct {
	ID              int
	OwnerID         int
	MenuName        string
	Deadline        time.Time
	PaymentReminder time.Time
	Status          int
}
