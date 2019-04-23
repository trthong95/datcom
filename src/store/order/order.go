package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type orderService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &orderService{
		db: db,
	}
}

func (os *orderService) Add(o *domain.OrderInput) (*models.Order, error) {
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

func (os *orderService) Delete(o *domain.OrderInput) error {
	order, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).One(context.Background(), os.db)
	if err != nil {
		return err
	}

	_, err = order.Delete(context.Background(), os.db)

	return err
}

func (os *orderService) Exist(o *domain.OrderInput) (bool, error) {
	b, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).Exists(context.Background(), os.db)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (os *orderService) Get(userID int) ([]*domain.Item, error) {
	orders, err := models.Orders(qm.Where("user_id=?", userID)).All(context.Background(), os.db)
	if err != nil {
		return nil, err
	}

	items := make([]*models.Item, len(orders))
	returnItems := make([]*domain.Item, len(orders))

	for i, order := range orders {
		items[i], err = models.Items(qm.Where("id=?", order.ItemID)).One(context.Background(), os.db)
		returnItems[i] = &domain.Item{
			ID:       items[i].ID,
			ItemName: items[i].ItemName,
			MenuID:   items[i].MenuID,
		}
	}

	return returnItems, err
}

// DeleteOrder ..
func (s *orderService) DeleteOrder(o *models.Order) error {
	_, err := o.Delete(context.Background(), s.db)
	return err
}

// CheckOrderExistByItemID ..
func (s *orderService) CheckOrderExistByItemID(ItemID int) (bool, error) {
	return models.Orders(
		qm.Where("item_id=?", ItemID),
	).Exists(context.Background(), s.db)
}

// GetAllOrdersByItemID ..
func (s *orderService) GetAllOrdersByItemID(ItemID int) ([]*models.Order, error) {
	return models.Orders(
		qm.Where("item_id=?", ItemID),
	).All(context.Background(), s.db)
}
