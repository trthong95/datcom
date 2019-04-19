package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/order"
)

//AddOrder adds new order to resource if order does not exist, otherwise return error
func (s *Service) AddOrder(o *order.Order) (*models.Order, error) {
	exist, _ := s.Order.Exist(o)
	if exist {
		return nil, errors.New("Order already exists")
	}

	return s.Order.Add(o)
}

//DeleteOrder deletes order from resource if order exists, otherwise return error
func (s *Service) DeleteOrder(o *order.Order) error {
	exist, _ := s.Order.Exist(o)
	if !exist {
		return errors.New("Order does not exist")
	}

	err := s.Order.Delete(o)

	return err
}

//GetOrdersByUserID gets items in user order by user ID
func (s *Service) GetOrdersByUserID(userID int) ([]*item.Item, error) {
	items, err := s.Order.Get(userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}
