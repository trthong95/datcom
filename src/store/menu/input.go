package menu

import "time"

// Menu ..
type Menu struct {
	ID              int
	OwnerID         int
	MenuName        string
	Deadline        time.Time
	PaymentReminder time.Time
	Status          int
}
