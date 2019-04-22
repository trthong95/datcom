package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (s *Service) AddPeopleInCharge(p *domain.PeopleInCharge) (*models.PeopleInCharge, error) {
	exist, _ := s.PeopleInCharge.Exist(p)
	if exist {
		return nil, errors.New("already exist")
	}

	return s.PeopleInCharge.Add(p)
}

func (s *Service) GetPeopleInCharge(menuID int) ([]*domain.PeopleInCharge, error) {
	people, err := s.PeopleInCharge.Get(menuID)
	if err != nil {
		return nil, err
	}

	return people, nil
}
