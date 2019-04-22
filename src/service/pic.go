package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (s *Service) AddPIC(p *domain.PICInput) (*models.PeopleInCharge, error) {
	exist, _ := s.Store.PIC.Exist(p)
	if exist {
		return nil, errors.New("already exist")
	}

	return s.Store.PIC.Add(p)
}

func (s *Service) GetPICByMenuID(menuID int) ([]*models.PeopleInCharge, error) {
	people, err := s.Store.PIC.GetByMenuID(menuID)
	if err != nil {
		return nil, err
	}

	return people, nil
}
