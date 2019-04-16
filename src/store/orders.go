package store

import (
	"context"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
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
		return errors.New("Order already exist")
	}
	return err
}

//Delete existing orders in database
func (store *storePostgres) DeleteOrders(userID int, itemID int) error {
	var order *models.Order
	// Select the order
	order, _ = models.Orders(qm.Where("user_id=? and item_id=?", userID, itemID)).One(context.Background(), store.db)
	if order == nil {
		// If not exists, do nothing
		return errors.New("Order does not exist")
	}
	_, err := order.Delete(context.Background(), store.db)

	return err
}

// Get items in user order
func (store *storePostgres) GetOrder(userID int) ([]*models.Item, error) {
	// Get orders of the user
	orders, err := models.Orders(qm.Where("user_id=?", userID)).All(context.Background(), store.db)
	items := make([]*models.Item, len(orders))
	for i, order := range orders {
		items[i], err = models.Items(qm.Where("id=?", order.ItemID)).One(context.Background(), store.db)
	}

	return items, err
}
