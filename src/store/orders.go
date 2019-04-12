package store

import (
	"context"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"datcom/backend/models"
)

// Inserts orders to database
func (store *storePostgres) AddOrders(userID int, itemID int) error {
	var order models.Order
	order.UserID = userID
	order.ItemID = itemID
	// First check if the order exists
	exists, err := models.Orders(qm.Where("user_id=? and item_id=?", userID, itemID)).Exists(context.Background(), store.db)
	if exists == false {
		// If not exist, insert new order
		err = order.Insert(context.Background(), store.db, boil.Infer())
	} else {
		// Else do nothing
		err = nil
	}
	return err
}

//Delete existing orders in database
func (store *storePostgres) DeleteOrders(userID int, itemID int) error {
	var order *models.Order
	// Select the order
	order, err := models.Orders(qm.Where("user_id=? and item_id=?", userID, itemID)).One(context.Background(), store.db)
	if &order == nil {
		// If not exists, do nothing
	} else {
		// If exists, delete the order
		_, err = order.Delete(context.Background(), store.db)
	}
	return err
}
