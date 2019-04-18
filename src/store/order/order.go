package order

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"datcom/backend/models"
	"datcom/backend/src/store/item"
)

type OrderService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &OrderService{
		db: db,
	}
}

func (os *OrderService) Add(o *Order) error {
	order := models.Order{
		UserID: o.UserID,
		ItemID: o.ItemID,
	}
	err := order.Insert(context.Background(), os.db, boil.Infer())
	return err
}

func (os *OrderService) Delete(o *Order) error {
	order, _ := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).One(context.Background(), os.db)
	_, err := order.Delete(context.Background(), os.db)
	return err
}

func (os *OrderService) Exist(o *Order) bool {
	err, _ := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).Exists(context.Background(), os.db)
	return err
}

func (os *OrderService) Get(userID int) ([]*item.Item, error) {
	// Select orders that match user_id
	orders, err := models.Orders(qm.Where("user_id=?", userID)).All(context.Background(), os.db)
	// Create needed lists
	items := make([]*models.Item, len(orders))
	returnItems := make([]*item.Item, len(orders))
	// Loop to get items list
	for i, order := range orders {
		items[i], err = models.Items(qm.Where("id=?", order.ItemID)).One(context.Background(), os.db)
		returnItems[i] = &item.Item{
			ID:       items[i].ID,
			ItemName: items[i].ItemName,
			MenuID:   items[i].MenuID,
		}
	}
	return returnItems, err
}
