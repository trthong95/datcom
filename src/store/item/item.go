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

func (s *itemService) FindByID(it *domain.Item) (*models.Item, error) {
	i := mapItemInputToModel(it)
	return models.FindItem(context.Background(), s.db, i.ID)
}

func (s *itemService) Delete(i *models.Item) error {
	_, err := i.Delete(context.Background(), s.db)
	return err
}

func (s *itemService) CheckItemExist(it *domain.Item) (bool, error) {
	i := mapItemInputToModel(it)
	return models.Items(
		qm.Where("(item_name = ? AND menu_id = ?) OR (id = ? AND menu_id = ?)", i.ItemName, i.MenuID, i.ID, i.MenuID),
	).Exists(context.Background(), s.db)
}
