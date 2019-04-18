package order

import "git.d.foundation/datcom/backend/src/store/item"

type Service interface {
	Add(o *Order) error
	Delete(o *Order) error
	Exist(o *Order) bool
	Get(userID int) ([]*item.Item, error)
}
