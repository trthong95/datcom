package store

import "github.com/hieunmce/datcom/models"

type Migrator interface {
	MigrateDB() error
	ReverseDB() error
	OneMigrateDB() error
}

type Store interface {
	GetAllUsers() ([]*models.User, error)
}
