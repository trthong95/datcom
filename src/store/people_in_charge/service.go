package people_in_charge

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(p *domain.PeopleInCharge) (*models.PeopleInCharge, error)
	Get(menuID int) ([]*domain.PeopleInCharge, error)
	Exist(p *domain.PeopleInCharge) (bool, error)
}
