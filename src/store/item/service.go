package item

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	Add(*domain.Item) (*models.Item, error)
	FindAnItem(i *Item) (*models.Item, error)
	DeleteAnItem(i *models.Item) (*models.Item, error)
	CheckItemExist(i *domain.Item) (bool, error)
}
