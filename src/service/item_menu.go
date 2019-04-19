package service

import (
	"database/sql"
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
)

// AddItems ..
func (s *Service) AddItems(items *item.Items, mn *menu.Menu) ([]*models.Item, error) {
	var list []*models.Item

	foundMenu, err := s.Menu.FindMenu(&menu.Menu{MenuName: mn.MenuName})
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("Non-exist-menu")
			return nil, err
		}
		return nil, err
	}

	for _, it := range items.ItemNames {
		i := &item.Item{ItemName: it.ItemName}
		m := &menu.Menu{ID: foundMenu.ID}
		exists, err := s.Item.CheckItemExist(i, m)
		if err != nil {
			return nil, err
		}
		if !exists {
			i.MenuID = foundMenu.ID
			resItem, err := s.Item.AddAnItem(i)
			if err != nil {
				return nil, err
			}
			list = append(list, resItem)
		}
	}

	return list, err
}
