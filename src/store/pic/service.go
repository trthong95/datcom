package pic

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(p *domain.PICInput) (*models.PeopleInCharge, error)
	GetByMenuID(menuID int) ([]*models.PeopleInCharge, error)
	Exist(p *domain.PICInput) (bool, error)
}
