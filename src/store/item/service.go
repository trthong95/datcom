package item

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	Add(i *domain.Item) (*models.Item, error)
	FindByID(itemID int) (*models.Item, error)
	Delete(i *models.Item) error
	CheckItemExist(it *domain.Item) (bool, error)
}
