package user

import (
	"context"
	"database/sql"
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type userService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &userService{
		db: db,
	}
}

func (us *userService) Create(p *domain.PersonInfo) (*models.User, error) {
	u := &models.User{Name: p.Name, Email: p.Email, Token: p.Token}
	err := u.Insert(context.Background(), us.db, boil.Infer())
	if err != nil {
		return nil, errors.New("Insert failed")
	}
	return us.Find(p)
}

func (us *userService) Find(p *domain.PersonInfo) (*models.User, error) {
	user, err := models.Users(qm.Where("email=?", p.Email)).One(context.Background(), us.db)
	return user, err
}

func (us *userService) FindAll() ([]*models.User, error) {
	users, err := models.Users().All(context.Background(), us.db)
	return users, err
}
