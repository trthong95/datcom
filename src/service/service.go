package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/people_in_charge"
	"git.d.foundation/datcom/backend/src/store/user"
)

type Service struct {
	User           user.Service
	Order          order.Service
	PeopleInCharge people_in_charge.Service
}

func NewService(db *sql.DB) *Service {
	return &Service{
		User:           user.NewService(db),
		Order:          order.NewService(db),
		PeopleInCharge: people_in_charge.NewService(db),
	}
}
