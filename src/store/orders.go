package store

import (
	"context"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"datcom/backend/models"
)

func (store *storePostgres) AddOrders(userID int, itemID int) error {
	var order models.Order
	var err error
	order.UserID = userID
	order.ItemID = itemID
	// First find the order if it exists
	result, err := models.Orders(qm.Where("user_id=? and item_id=?", userID, itemID)).One(context.Background(), store.db)
	if result == nil {
		// If not exist, insert new order
		err = order.Insert(context.Background(), store.db, boil.Infer())
	} else {
		// Else do nothing
		err = nil
	}
	return err
}
