package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//AddOrder adds new order to resource if order does not exist, otherwise return error
func (s *Service) AddOrder(o *domain.OrderInput) (*models.Order, error) {
	exist, _ := s.Store.Order.Exist(o)
	if exist {
		return nil, errors.New("Order already exists")
	}

	return s.Store.Order.Add(o)
}

//DeleteOrder deletes order from resource if order exists, otherwise return error
func (s *Service) DeleteOrder(o *domain.OrderInput) error {
	exist, _ := s.Store.Order.Exist(o)
	if !exist {
		return errors.New("Order does not exist")
	}

	err := s.Store.Order.Delete(o)

	return err
}

//GetOrdersByUserID gets items in user order by user ID
func (s *Service) GetOrdersByUserID(userID int) ([]*domain.Item, error) {
	items, err := s.Store.Order.Get(userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}
