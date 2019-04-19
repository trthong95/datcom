package order

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
)

type Service interface {
	Add(o *Order) (*models.Order, error)
	Delete(o *Order) error
	Exist(o *Order) (bool, error)
	Get(userID int) ([]*item.Item, error)
}
