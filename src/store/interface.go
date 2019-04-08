package store

import "github.com/hieunmce/datcom/models"

type Migrator interface {
	MigrateDB() error
	ReverseDB() error
	MigrateOneDB() error
	ReverseOneDB() error
}

type Store interface {
	GetAllUsers() ([]*models.User, error)
	SetUserAndEmail(name string, email string) error
}
