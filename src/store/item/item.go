package item

import (
	"context"
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/menu"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

func (s *itemService) AddAnItem(i *Item) (*models.Item, error) {
	item := &models.Item{ItemName: i.ItemName, MenuID: i.MenuID}
	err := item.Insert(context.Background(), s.db, boil.Infer())
	return item, err
}

func (s *itemService) CheckItemExist(i *Item, m *menu.Menu) (bool, error) {
	exists, err := models.Items(
		qm.Where("(item_name = ? AND menu_id = ?) OR (id = ? AND menu_id = ?)",
			i.ItemName,
			m.ID,
			i.ID,
			m.ID)).Exists(context.Background(), s.db)
	return exists, err
}
