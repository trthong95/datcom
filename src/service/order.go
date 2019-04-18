package service

import (
	"datcom/backend/src/store/item"
	"datcom/backend/src/store/order"
	"errors"
)

//AddOrder Add order to resource
func (s *Service) AddOrder(o *order.Order) error {
	// First check if the order exists
	exist := s.Order.Exist(o)
	if exist == true {
		// If exists return error
		return errors.New("Order already exists")
	}
	// If exist, add order
	err := s.Order.Add(o)
	return err
}

//DeleteOrder Delete existing orders in resource
func (s *Service) DeleteOrder(o *order.Order) error {
	// Check if order exists
	exist := s.Order.Exist(o)
	if exist == false {
		// If not exists, return error
		return errors.New("Order does not exist")
	}
	// If exist delete the order
	err := s.Order.Delete(o)
	return err
}

//GetOrders Get items in user order
func (s *Service) GetOrders(userID int) ([]*item.Item, error) {
	// Get orders of the user
	items, err := s.Order.Get(userID)
	return items, err
}
