package item

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type itemService struct {
	db *sql.DB
}

// NewService ..
func NewService(db *sql.DB) Service {
	return &itemService{
		db: db,
	}
}

func mapItemInputToModel(i *domain.Item) *models.Item {
	return &models.Item{
		ID:       i.ID,
		ItemName: i.ItemName,
		MenuID:   i.MenuID,
	}
}

func (s *itemService) Add(i *domain.Item) (*models.Item, error) {
	item := mapItemInputToModel(i)
	return item, item.Insert(context.Background(), s.db, boil.Infer())
}

func (s *itemService) CheckItemExist(i *domain.Item) (bool, error) {
	item := mapItemInputToModel(i)
	return models.Items(
		qm.Where("(item_name = ? AND menu_id = ?) OR (id = ? AND menu_id = ?)", item.ItemName, item.MenuID, item.ID, item.MenuID),
	).Exists(context.Background(), s.db)
}
