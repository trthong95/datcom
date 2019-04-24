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

func mapMenuInputToModel(mn *domain.MenuInput) *models.Menu {
	return &models.Menu{
		ID:              mn.ID,
		OwnerID:         mn.OwnerID,
		MenuName:        mn.MenuName,
		Deadline:        mn.Deadline,
		PaymentReminder: mn.PaymentReminder,
		Status:          mn.Status,
	}
}

func (s *menuService) CheckMenuExist(menuID int) (bool, error) {
	return models.Menus(
		qm.Where("id = ?", menuID),
	).Exists(context.Background(), s.db)
}
