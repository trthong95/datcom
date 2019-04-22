package item

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	Add(*domain.Item) (*models.Item, error)
	FindByID(*domain.Item) (*models.Item, error)
	Delete(i *models.Item) error
	CheckItemExist(it *domain.Item) (bool, error)
}
