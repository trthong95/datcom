package menu

import (
	"context"
	"database/sql"

	"git.d.foundation/datcom/backend/models"

	"github.com/volatiletech/sqlboiler/queries/qm"
)

type menuService struct {
	db *sql.DB
}

// NewService ..
func NewService(db *sql.DB) Service {
	return &menuService{
		db: db,
	}
}

func (s *menuService) FindMenu(m *Menu) (*models.Menu, error) {
	menu, err := models.Menus(qm.Where("id = ? OR menu_name = ?", m.ID, m.MenuName)).One(context.Background(), s.db)
	return menu, err
}
