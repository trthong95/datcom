package item

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/menu"
)

// Service ..
type Service interface {
	AddAnItem(i *Item) (*models.Item, error)
	CheckItemExist(i *Item, m *menu.Menu) (bool, error)
}
