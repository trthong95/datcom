package order

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(o *domain.OrderInput) (*models.Order, error)
	Delete(o *domain.OrderInput) error
	Exist(o *domain.OrderInput) (bool, error)
	Get(userID int) ([]*domain.Item, error)
	DeleteO(o *models.Order) error
	CheckOrderExistByItemID(ItemID int) (bool, error)
	GetAllOrdersByItemID(ItemID int) ([]*models.Order, error)
}
