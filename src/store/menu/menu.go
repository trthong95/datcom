package menu

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
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

func mapMenuInputToModel(m *domain.MenuInput) *models.Menu {
	return &models.Menu{
		ID:              m.ID,
		OwnerID:         m.OwnerID,
		MenuName:        m.MenuName,
		Deadline:        m.Deadline,
		PaymentReminder: m.PaymentReminder,
		Status:          m.Status,
	}
}

func (s *menuService) CheckMenuExist(menuID int) (bool, error) {
	return models.Menus(
		qm.Where("id = ?", menuID),
	).Exists(context.Background(), s.db)
}

func (s *menuService) FindByID(mn *domain.MenuInput) (*models.Menu, error) {
	m := mapMenuInputToModel(mn)
	return models.FindMenu(context.Background(), s.db, m.ID)
}
