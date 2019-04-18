package store

type Migrator interface {
	MigrateDB() error
	ReverseDB() error
}
