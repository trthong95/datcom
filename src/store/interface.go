package store

import "gitlab.com/thongnt/lunch-order/models"

type Migrator interface {
	MigrateDB() error
	ReverseDB() error
}

type Store interface {
	GetAllUsers() ([]*models.User, error)
}
