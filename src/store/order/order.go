package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
)

type OrderService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &OrderService{
		db: db,
	}
}

func (os *OrderService) Add(o *Order) (*models.Order, error) {
	order := &models.Order{
		UserID: o.UserID,
		ItemID: o.ItemID,
	}

	err := order.Insert(context.Background(), os.db, boil.Infer())
	if err != nil {
		return nil, errors.New("Cannot insert")
	}

	return order, err
}

func (os *OrderService) Delete(o *Order) error {
	order, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).One(context.Background(), os.db)
	if err != nil {
		return err
	}

	_, err = order.Delete(context.Background(), os.db)

	return err
}

func (os *OrderService) Exist(o *Order) (bool, error) {
	b, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).Exists(context.Background(), os.db)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (os *OrderService) Get(userID int) ([]*item.Item, error) {
	orders, err := models.Orders(qm.Where("user_id=?", userID)).All(context.Background(), os.db)
	if err != nil {
		return nil, err
	}

	items := make([]*models.Item, len(orders))
	returnItems := make([]*item.Item, len(orders))

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
