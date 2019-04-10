package store

import "datcom/backend/models"

type Migrator interface {
	MigrateDB() error
	ReverseDB() error
}

type Store interface {
	GetAllUsers() ([]*models.User, error)
}
