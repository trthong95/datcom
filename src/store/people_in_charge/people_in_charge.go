package people_in_charge

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"github.com/volatiletech/sqlboiler/boil"
)

type PeopleInChargeService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &PeopleInChargeService{
		db: db,
	}
}

func (pics *PeopleInChargeService) Add(p *domain.PeopleInCharge) (*models.PeopleInCharge, error) {
	people := &models.PeopleInCharge{
		UserID: p.UserID,
		MenuID: p.MenuID,
	}

	err := people.Insert(context.Background(), pics.db, boil.Infer())
	if err != nil {
		return nil, errors.New("cannot insert")
	}

	return people, nil
}

func (pics *PeopleInChargeService) Get(menuID int) ([]*domain.PeopleInCharge, error) {
	people, err := models.PeopleInCharges(qm.Where("menu_id=?", menuID)).All(context.Background(), pics.db)
	if err != nil {
		return nil, err
	}

	returnPeople := make([]*domain.PeopleInCharge, len(people))

	for i, person := range people {
		returnPeople[i] = &domain.PeopleInCharge{
			UserID: person.UserID,
			MenuID: person.MenuID,
		}
	}

	return returnPeople, nil
}

func (pics *PeopleInChargeService) Exist(p *domain.PeopleInCharge) (bool, error) {
	b, err := models.PeopleInCharges(qm.Where("user_id=? and menu_id=?", p.UserID, p.MenuID)).Exists(context.Background(), pics.db)
	if err != nil {
		return false, err
	}

	return b, nil
}
