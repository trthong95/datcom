package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// DeleteItem ..
func (s *Service) DeleteItem(i *domain.Item) (*models.Item, error) {

	exists, err := s.Store.Item.CheckItemExist(&domain.Item{ID: i.ID, MenuID: i.MenuID})
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Item does not exist")
	}

	it, err := s.Store.Item.FindByID(i.ID)
	if err != nil {
		return nil, err
	}

	exists, err = s.Store.Order.CheckOrderExistByItemID(i.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		orders, err := s.Store.Order.GetAllOrdersByItemID(i.ID)
		if err != nil {
			return nil, err
		}
		for _, o := range orders {
			err := s.Store.Order.DeleteOrder(o)
			if err != nil {
				return nil, err
			}
		}
	}

	err = s.Store.Item.Delete(it)
	if err != nil {
		return nil, err
	}

	return it, err
}
