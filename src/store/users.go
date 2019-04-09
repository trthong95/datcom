package store

import (
	"context"

	"datcom/backend/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type PersonInfo struct {
	Name  string
	Email string
	Token string
}

func (store *storePostgres) GetAllUsers() ([]*models.User, error) {
	users, err := models.Users().All(context.Background(), store.db)
	return users, err
}

func (store *storePostgres) CreateUser(p *PersonInfo) error {
	u := &models.User{Name: p.Name, Email: p.Email, Token: p.Token}
	err := u.Upsert(context.Background(), store.db, true, []string{"email"}, boil.Whitelist("name"), boil.Infer())
	return err
}
