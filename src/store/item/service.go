package item

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	Add(*domain.Item) (*models.Item, error)
	CheckItemExist(*domain.Item) (bool, error)
}
