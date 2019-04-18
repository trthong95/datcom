package user

import (
	"context"
	"database/sql"
	"datcom/backend/models"
	"datcom/backend/src/domain"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &UserService{
		db: db,
	}
}

func (us *UserService) Create(p *domain.PersonInfo) (*models.User, error) {
	u := &models.User{Name: p.Name, Email: p.Email, Token: p.Token}
	err := u.Insert(context.Background(), us.db, boil.Infer())
	if err != nil {
		return &models.User{}, errors.New("Insert faild")
	}
	return us.Find(p)
}

func (us *UserService) Find(p *domain.PersonInfo) (*models.User, error) {
	user, _ := models.Users(qm.Where("email=?", p.Email)).One(context.Background(), us.db)
	return user, nil
}

func (us *UserService) FindAll() ([]*models.User, error) {
	users, err := models.Users().All(context.Background(), us.db)
	return users, err
}
