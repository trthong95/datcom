package store

import (
	"context"

	"github.com/hieunmce/datcom/models"
)

func (store *storePostgres) GetAllUsers() ([]*models.User, error) {
	models.Users().All(context.Background(), store.db)
	return nil, nil
}
