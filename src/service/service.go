package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/pic"
	"git.d.foundation/datcom/backend/src/store/user"
)

type Service struct {
	Store Store
}

type Store struct {
	User  user.Service
	Order order.Service
	PIC   pic.Service
}

func NewService(db *sql.DB) *Service {
	return &Service{
		Store: Store{
			User:  user.NewService(db),
			Order: order.NewService(db),
			PIC:   pic.NewService(db),
		},
	}
}
