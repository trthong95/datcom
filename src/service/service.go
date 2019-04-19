package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/store/user"
)

type Service struct {
	User user.Service
}

func NewService(db *sql.DB) *Service {
	return &Service{
		User: user.NewService(db),
	}
}
