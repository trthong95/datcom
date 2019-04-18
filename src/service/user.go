package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (s *Service) CreateUser(p *domain.PersonInfo) (*models.User, error) {
	u, err := s.User.Find(p)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return s.User.Create(p)
	}
	return &models.User{}, errors.New("User existed")
}

func (s *Service) GetAllUser() ([]*models.User, error) {
	u, err := s.User.FindAll()
	return u, err
}
