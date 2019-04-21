package menu

import (
	"context"
	"database/sql"

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

func (s *menuService) FindByID(mn *domain.MenuInput) (*models.Menu, error) {
	m := mapMenuInputToModel(mn)
	return models.FindMenu(context.Background(), s.db, m.ID)
}
