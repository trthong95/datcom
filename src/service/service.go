package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/pic"
	"git.d.foundation/datcom/backend/src/store/user"
)

type Service struct {
	Store Store
}

type Store struct {
	User  user.Service
	Item  item.Service
	Menu  menu.Service
	Order order.Service
	PIC   pic.Service
}

func NewService(db *sql.DB) *Service {
	return &Service{
		Store: Store{
			User:  user.NewService(db),
			Item:  item.NewService(db),
			Menu:  menu.NewService(db),
			Order: order.NewService(db),
			PIC:   pic.NewService(db),
		},
	}
}
