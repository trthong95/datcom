package store

import (
	"context"

	"datcom/backend/models"
)

func (store *storePostgres) GetAllUsers() ([]*models.User, error) {
	models.Users().All(context.Background(), store.db)
	return nil, nil
}
