package store

import (
	"datcom/backend/models"
)

// Migrator ...
type Migrator interface {
	MigrateDB() error
	ReverseDB() error
}

// Store ...
type Store interface {
	GetAllUsers() ([]*models.User, error)
	CreateUser(p *PersonInfo) error
	InsertItemsToDB(itemNames []string, menuName string) error
	AddOrders(userID int, itemID int) error
}
