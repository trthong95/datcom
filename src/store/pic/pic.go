package pic

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type picService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &picService{
		db: db,
	}
}

func (pics *picService) Add(p *domain.PICInput) (*models.PeopleInCharge, error) {
	people := &models.PeopleInCharge{
		UserID: p.UserID,
		MenuID: p.MenuID,
	}

	err := people.Insert(context.Background(), pics.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return people, nil
}

func (pics *picService) GetByMenuID(menuID int) ([]*models.PeopleInCharge, error) {
	people, err := models.PeopleInCharges(qm.Where("menu_id=?", menuID)).All(context.Background(), pics.db)
	if err != nil {
		return nil, err
	}

	returnPeople := make([]*models.PeopleInCharge, len(people))

	for i, person := range people {
		returnPeople[i] = &models.PeopleInCharge{
			UserID: person.UserID,
			MenuID: person.MenuID,
		}
	}

	return returnPeople, nil
}

func (pics *picService) Exist(p *domain.PICInput) (bool, error) {
	b, err := models.PeopleInCharges(qm.Where("user_id=? AND menu_id=?", p.UserID, p.MenuID)).Exists(context.Background(), pics.db)
	if err != nil {
		return false, err
	}

	return b, nil
}
