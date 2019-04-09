package store

import (
	"context"

	"github.com/trthong95/datcom/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func (store *storePostgres) GetAllUsers() ([]*models.User, error) {
	user, err := models.Users().All(context.Background(), store.db)
	return user, err
}

func (store *storePostgres) SetUserAndEmail(name string, email string) error {
	u := &models.User{}
	u.Name = name
	u.Email = email
	err := u.Upsert(context.Background(), store.db, true, []string{"email"}, boil.Whitelist("name"), boil.Infer())
	return err
}

// func (store *storePostgres) UpdateEmail(name)
