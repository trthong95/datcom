package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// AddItems ..
func (s *Service) AddItems(items *domain.ItemInput, menuID int) ([]*models.Item, error) {
	var list []*models.Item

	exists, err := s.Store.Menu.CheckMenuExist(menuID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Menu does not exist")
	}

	for _, it := range items.Items {
		i := &domain.Item{ItemName: it.ItemName, MenuID: menuID}
		exists, err := s.Store.Item.CheckItemExist(i)
		if err != nil {
			return nil, err
		}
		if exists {
			continue
		}
		resItem, err := s.Store.Item.Add(i)
		if err != nil {
			return nil, err
		}
		list = append(list, resItem)
	}

	return list, err
}
