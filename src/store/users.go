package store

import (
	"context"

	"gitlab.com/thongnt/lunch-order/models"
)

func (store *storePostgres) GetAllUsers() ([]*models.User, error) {
	models.Users().All(context.Background(), store.db)
	return nil, nil
}
