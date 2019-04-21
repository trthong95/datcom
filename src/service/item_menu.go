package service

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// AddItems ..
func (s *Service) AddItems(items *domain.ItemInput, mn *domain.MenuInput) ([]*models.Item, error) {
	var list []*models.Item

	m, err := s.Store.Menu.FindByID(&domain.MenuInput{ID: mn.ID})
	if err != nil {
		return nil, err
	}

	for _, it := range items.Items {
		i := &domain.Item{ItemName: it.ItemName, MenuID: m.ID}
		exists, err := s.Store.Item.CheckItemExist(i)
		if err != nil {
			return nil, err
		}
		if exists {
			continue
		}
		i.MenuID = m.ID
		resItem, err := s.Store.Item.Add(i)
		if err != nil {
			return nil, err
		}
		list = append(list, resItem)
	}

	return list, err
}
